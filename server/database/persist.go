package database

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"

	"github.com/alexflint/go-memdump"
)

// LoadData loads the persisted data into the index tree.
func (db *DB) LoadData() {
	db.ioRWMutex.RLock()
	defer db.ioRWMutex.RUnlock()

	f, err := db.path.OpenFile("data")
	if err != nil {
		log.Fatalln(err)
		return
	}

	r := bufio.NewReader(f)

	var data *KeyValuePair
	if err := memdump.Decode(r, &data); err != nil {
		log.Fatalln(err)
		return
	}

	if err := db.index.Insert(int(data.Hash), data.JSON); err != nil {
		return
	}
}

// WriteDisk writes a key value pair to the disk for persistence.
func (db *DB) WriteDisk(hash uint32, value interface{}) {
	db.ioRWMutex.Lock()
	defer db.ioRWMutex.Unlock()

	if err := db.path.Mkdir(); err != nil {
		fmt.Println(err)
		return
	}

	f, err := db.path.OpenFile("data")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	m, err := json.Marshal(value)
	if err != nil {
		return
	}

	data := KeyValuePair{
		Hash: hash,
		JSON: m,
	}

	if err := memdump.Encode(w, &data); err != nil {
		fmt.Println(err)
		return
	}
}

// KeyValuePair is a serialized key value pair to the persistence file.
type KeyValuePair struct {
	Hash uint32
	JSON []byte
}
