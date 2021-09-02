package utxo

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/piotrnar/gocoin/lib/btc"
)

type UTXOSqlite3 struct {
	UtxoDBNoops
	UtxoEmptyLocks
	UtxoEmptyCallbacks
	UtxoNotCompressed

	filename string
	db       *sql.DB
}

func errHandler(err error) {
	panic(err)
}

func NewUnspentDb_sqlite3(opts *NewUnspentOpts) *UTXOSqlite3 {

	us3 := &UTXOSqlite3{
		filename: opts.Dir,
	}

	db, err := sql.Open("sqlite3", us3.filename)
	if err != nil {
		errHandler(err)
	}
	us3.db = db

	if err := us3.init(); err != nil {
		errHandler(err)
	}

	return us3
}

func (db *UTXOSqlite3) init() error {

	create := `CREATE TABLE IF NOT EXISTS hashes (id INTEGER PRIMARY KEY,
		hash BLOB,
		value BLOB);`

	_, err := db.db.Exec(create)

	return err
}

func (db *UTXOSqlite3) HashMap() map[UtxoKeyType][]byte {
	panic("not implemented") // TODO: Implement
}

func (db *UTXOSqlite3) HashMapIdx(ind UtxoKeyType) []byte {
	var value []byte

	row := db.db.QueryRow("SELECT value FROM hashes WHERE hash=?", ind)
	err := row.Scan(value)
	if err != nil && err != sql.ErrNoRows {
		errHandler(err)
	}

	return value
}

func (db *UTXOSqlite3) HashMapSetIdx(ind UtxoKeyType, value []byte) {
	panic("not implemented") // TODO: Implement
}

func (db *UTXOSqlite3) HashMapLen() int {
	var count int

	row := db.db.QueryRow("SELECT COUNT(*) FROM hashes")
	err := row.Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		errHandler(err)
	}

	return count
}

func (db *UTXOSqlite3) WritingInProgress() bool {
	// Handled by the sqlite3 engine
	return false
}

func (db *UTXOSqlite3) CurrentHeightOnDisk() uint32 {
	panic("not implemented") // TODO: Implement
}

func (db *UTXOSqlite3) LastBlockHash() []byte {
	panic("not implemented") // TODO: Implement
}

func (db *UTXOSqlite3) SetLastBlockHeight(v uint32) {
	panic("not implemented") // TODO: Implement
}

func (db *UTXOSqlite3) LastBlockHeight() uint32 {
	panic("not implemented") // TODO: Implement
}

// CommitBlockTxs commits the given add/del transactions to UTXO and Unwind DBs.
func (db *UTXOSqlite3) CommitBlockTxs(_ *BlockChanges, _ []byte) error {
	panic("not implemented") // TODO: Implement
}

func (db *UTXOSqlite3) UndoBlockTxs(_ *btc.Block, _ []byte) {
	panic("not implemented") // TODO: Implement
}

func (db *UTXOSqlite3) Save() bool {
	panic("not implemented") // TODO: Implement
}

// Close flushes the data and closes all the files.
func (db *UTXOSqlite3) Close() {
	panic("not implemented") // TODO: Implement
}

// UnspentGet gets the given unspent output.
func (db *UTXOSqlite3) UnspentGet(_ *btc.TxPrevOut) *btc.TxOut {
	panic("not implemented") // TODO: Implement
}

// TxPresent returns true if gived TXID is in UTXO.
func (db *UTXOSqlite3) TxPresent(_ *btc.Uint256) bool {
	panic("not implemented") // TODO: Implement
}

func (db *UTXOSqlite3) AbortWriting() {
	panic("not implemented") // TODO: Implement
}

func (db *UTXOSqlite3) UTXOStats() string {
	panic("not implemented") // TODO: Implement
}

// GetStats returns DB statistics.
func (db *UTXOSqlite3) GetStats() string {
	panic("not implemented") // TODO: Implement
}

func (db *UTXOSqlite3) PurgeUnspendable(_ bool) {
	panic("not implemented") // TODO: Implement
}
