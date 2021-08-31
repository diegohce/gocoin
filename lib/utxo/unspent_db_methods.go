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
