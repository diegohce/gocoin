package utxo

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/piotrnar/gocoin/lib/btc"
)

type UTXOSqlite3 struct {
	UtxoDBNoops
	UtxoEmptyLocks
	UtxoEmptyCallbacks
	UtxoNotCompressed
	UtxoLastBlockHeight

	filename string
	db       *sql.DB
	log      *log.Logger
}

func NewUnspentDb_sqlite3(opts *NewUnspentOpts) *UTXOSqlite3 {

	us3 := &UTXOSqlite3{
		filename: opts.Dir,
		log:      log.New(os.Stdout, "[SQLite3 backend] ", log.Default().Flags()),
	}

	db, err := sql.Open("sqlite3", us3.filename)
	if err != nil {
		us3.errHandler(err)
	}
	us3.db = db

	if err := us3.init(); err != nil {
		us3.errHandler(err)
	}

	return us3
}

func (db *UTXOSqlite3) errHandler(err error) {
	db.log.Panic(err)
}

func (db *UTXOSqlite3) init() error {

	create := `CREATE TABLE IF NOT EXISTS hashes (id INTEGER PRIMARY KEY,
		hash BLOB NOT NULL UNIQUE,
		value BLOB NOT NULL);`

	//create unique index indx_email on emp (email);

	_, err := db.db.Exec(create)

	return err
}

func (db *UTXOSqlite3) HashMap() map[UtxoKeyType][]byte {

	rows, err := db.db.Query("SELECT hash, value FROM hashes")
	if err != nil {
		db.errHandler(err)
	}
	defer rows.Close()

	hmapLen := db.HashMapLen()
	hmap := make(map[UtxoKeyType][]byte, hmapLen)

	for rows.Next() {
		var hash [UtxoIdxLen]byte
		var value []byte

		if err := rows.Scan(&hash, &value); err != nil {
			if err == sql.ErrNoRows {
				db.log.Println("HashMap() no rows")
				return hmap
			}
			db.errHandler(err)
		}
		hmap[hash] = value
	}

	return hmap
}

func (db *UTXOSqlite3) HashMapIdx(ind UtxoKeyType) []byte {
	var value []byte

	row := db.db.QueryRow("SELECT value FROM hashes WHERE hash=?", ind)
	err := row.Scan(&value)
	if err != nil && err != sql.ErrNoRows {
		db.errHandler(err)
	}
	if err == sql.ErrNoRows {
		db.log.Println("HashMapIdx() no rows")
	}

	return value
}

func (db *UTXOSqlite3) HashMapSetIdx(ind UtxoKeyType, value []byte) {
	var ra int64
	var resErr error

	upsert := `INSERT INTO hashes (hash, value) VALUES (?, ?)
		ON CONFLICT(hash) DO
		UPDATE SET value=excluded.value`

	r, err := db.db.Exec(upsert, ind, value)
	if err != nil {
		db.errHandler(err)
	}

	ra, resErr = r.RowsAffected()
	db.log.Println("HashMapSetIdx(): rows affected", ra, "err", resErr)

}

func (db *UTXOSqlite3) HashMapLen() int {
	var count int

	row := db.db.QueryRow("SELECT COUNT(*) FROM hashes")
	err := row.Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		db.errHandler(err)
	}

	return count
}

func (db *UTXOSqlite3) WritingInProgress() bool {
	// Handled by the sqlite3 engine
	return false
}

func (db *UTXOSqlite3) CurrentHeightOnDisk() uint32 {
	return db.lastBlockHeight
}

func (db *UTXOSqlite3) LastBlockHash() []byte {
	var maxRowID int64
	var lastHash [UtxoIdxLen]byte

	sqlMaxRowID := `SELECT MAX(id) FROM hashes`
	sqlLastHash := `SELECT hash FROM hashes WHERE id=?`

	row := db.db.QueryRow(sqlMaxRowID)
	if err := row.Scan(&maxRowID); err != nil {
		db.log.Println("LastBlockHash() maxRowID:", err)
		return lastHash[:]
	}

	row = db.db.QueryRow(sqlLastHash, maxRowID)
	if err := row.Scan(&lastHash); err != nil {
		db.log.Println("LastBlockHash() lastHash:", err)
		return lastHash[:]
	}
	return lastHash[:]
}

// CommitBlockTxs commits the given add/del transactions to UTXO and Unwind DBs.
func (db *UTXOSqlite3) CommitBlockTxs(_ *BlockChanges, _ []byte) error {
	panic("not implemented") // TODO: Implement
}

func (db *UTXOSqlite3) UndoBlockTxs(_ *btc.Block, _ []byte) {
	panic("not implemented") // TODO: Implement
}

func (db *UTXOSqlite3) Save() bool {
	return true
}

// Close flushes the data and closes all the files.
func (db *UTXOSqlite3) Close() {
	//db.db.Close()
}

// UnspentGet gets the given unspent output.
func (db *UTXOSqlite3) UnspentGet(_ *btc.TxPrevOut) *btc.TxOut {
	panic("not implemented") // TODO: Implement
}

// TxPresent returns true if gived TXID is in UTXO.
func (db *UTXOSqlite3) TxPresent(id *btc.Uint256) bool {
	var ind UtxoKeyType

	copy(ind[:], id.Hash[:])

	sqlPresent := `SELECT hash FROM hashes WHERE hash=?`

	row := db.db.QueryRow(sqlPresent, ind)
	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			return false
		} else {
			db.log.Panic(row.Err())
		}
	}
	return true
}

func (db *UTXOSqlite3) AbortWriting() {}

func (db *UTXOSqlite3) UTXOStats() string {
	return "UTXOStats() not implemented"
}

// GetStats returns DB statistics.
func (db *UTXOSqlite3) GetStats() string {
	return "GetStats() not implemented"
}

func (db *UTXOSqlite3) PurgeUnspendable(_ bool) {

	res, err := db.db.Exec(`DELETE FROM hashes`)
	if err != nil {
		db.log.Println("PurgeUnspendable():", err)
	}

	ra, _ := res.RowsAffected()
	db.log.Println("purged", ra, "records")
}
