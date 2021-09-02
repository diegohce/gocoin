package utxo

import (
	"github.com/piotrnar/gocoin/lib/btc"
)

type UTXODB interface {
	HashMap() map[UtxoKeyType][]byte

	HashMapIdx(ind UtxoKeyType) []byte

	HashMapSetIdx(ind UtxoKeyType, value []byte)

	HashMapLen() int

	RWMutexRUnlock()

	RWMutexRLock()

	WritingInProgress() bool

	CurrentHeightOnDisk() uint32

	LastBlockHash() []byte

	SetLastBlockHeight(v uint32)

	LastBlockHeight() uint32

	DirtyDB() bool

	SetDirtyDB()

	UnwindBufLen() uint32

	ComprssedUTXO() bool

	SetComprssedUTXO(bool)

	SetCBNotifyTxAdd(func(*UtxoRec))

	SetCBNotifyTxDel(func(*UtxoRec, []bool))

	// CommitBlockTxs commits the given add/del transactions to UTXO and Unwind DBs.
	CommitBlockTxs(*BlockChanges, []byte) error

	UndoBlockTxs(*btc.Block, []byte)

	// Idle should be called when the main thread is idle.
	Idle() bool

	Save() bool

	HurryUp()

	// Close flushes the data and closes all the files.
	Close()

	// UnspentGet gets the given unspent output.
	UnspentGet(*btc.TxPrevOut) *btc.TxOut

	// TxPresent returns true if gived TXID is in UTXO.
	TxPresent(*btc.Uint256) bool

	AbortWriting()

	UTXOStats() string

	// GetStats returns DB statistics.
	GetStats() string

	PurgeUnspendable(bool)
}

type UtxoDBNoops struct{}

func (UtxoDBNoops) Idle() bool           { return true }
func (UtxoDBNoops) HurryUp()             {}
func (UtxoDBNoops) DirtyDB() bool        { return false }
func (UtxoDBNoops) SetDirtyDB()          {}
func (UtxoDBNoops) UnwindBufLen() uint32 { return 256 }

type UtxoEmptyLocks struct{}

func (UtxoEmptyLocks) RWMutexRUnlock() {}
func (UtxoEmptyLocks) RWMutexRLock()   {}

type UtxoEmptyCallbacks struct{}

func (UtxoEmptyCallbacks) SetCBNotifyTxAdd(fn func(*UtxoRec)) {}

func (UtxoEmptyCallbacks) SetCBNotifyTxDel(fn func(*UtxoRec, []bool)) {}

type UtxoNotCompressed struct{}

func (UtxoNotCompressed) ComprssedUTXO() bool     { return false }
func (UtxoNotCompressed) SetComprssedUTXO(_ bool) {}

type UtxoLastBlockHeight struct {
	lastBlockHeight uint32
}

func (u *UtxoLastBlockHeight) SetLastBlockHeight(v uint32) {
	u.lastBlockHeight = v
}

func (u *UtxoLastBlockHeight) LastBlockHeight() uint32 {
	return u.lastBlockHeight
}
