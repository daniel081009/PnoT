package db

import (
	"PnoT/config"

	"github.com/boltdb/bolt"
)

var (
	MyDB *bolt.DB
)

func init() {
	db, err := bolt.Open(config.C.DB_Path, 0600, nil)
	if err != nil {
		panic(err)
	}
	MyDB = db
}
