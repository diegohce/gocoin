package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/piotrnar/gocoin/lib/others/sys"
	"github.com/piotrnar/gocoin/lib/utxo"
)

func main() {
	var dir = ""

	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	if dir != "" && !strings.HasSuffix(dir, string(os.PathSeparator)) {
		dir += string(os.PathSeparator)
	}

	sys.LockDatabaseDir(dir)
	defer sys.UnlockDatabaseDir()

	db := utxo.NewUnspentDb("builtin", &utxo.NewUnspentOpts{Dir: dir})
	if db == nil {
		println("UTXO.db not found.")
		return
	}

	if !db.ComprssedUTXO() {
		fmt.Println("UTXO.db is not compressed.")
		return
	}

	fmt.Println("Decompressing UTXO records")
	for k, v := range db.HashMap() {
		rec := utxo.NewUtxoRecStatic(k, v)
		db.HashMapSetIdx(k, utxo.SerializeU(rec, false, nil))
	}
	db.SetComprssedUTXO(false)
	db.SetDirtyDB()
	fmt.Println("Saving new UTXO.db")
	db.Close()
	fmt.Println("Done")
	fmt.Println("WARNING: the undo folder has not been converted")
}
