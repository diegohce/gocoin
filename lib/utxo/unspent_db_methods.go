package utxo

func (db *UnspentDB) HashMap() map[UtxoKeyType][]byte {
	return db.hashMap
}

func (db *UnspentDB) HashMapIdx(ind UtxoKeyType) []byte {
	return db.hashMap[ind]
}

func (db *UnspentDB) HashMapSetIdx(ind UtxoKeyType, value []byte) {
	db.hashMap[ind] = value
}

func (db *UnspentDB) HashMapLen() int {
	return len(db.hashMap)
}

func (db *UnspentDB) LastBlockHash() []byte {
	return db.lastBlockHash
}

func (db *UnspentDB) SetLastBlockHeight(v uint32) {
	db.lastBlockHeight = v
}

func (db *UnspentDB) LastBlockHeight() uint32 {
	return db.lastBlockHeight
}

func (db *UnspentDB) DirtyDB() bool {
	return db.dirtyDB.Get()
}

func (db *UnspentDB) SetDirtyDB() {
	db.dirtyDB.Set()
}

func (db *UnspentDB) ComprssedUTXO() bool {
	return db.comprssedUTXO
}
func (db *UnspentDB) SetComprssedUTXO(v bool) {
	db.comprssedUTXO = v
}

func (db *UnspentDB) UnwindBufLen() uint32 {
	return db.unwindBufLen
}
