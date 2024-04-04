package db_auth

import (
	"PnoT/db"
	"fmt"

	"github.com/boltdb/bolt"
)

func CreateUser(username string, password string) error {
	return db.MyDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return err
		}
		bucket := tx.Bucket([]byte("users"))
		return bucket.Put([]byte(username), []byte(password))
	})
}
func LoginUser(username string, password string) error {
	return db.MyDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("users"))
		if bucket == nil {
			return fmt.Errorf("Invalid Username")
		}
		if string(bucket.Get([]byte(username))) == password {
			return nil
		}
		return fmt.Errorf("Invalid Password")
	})
}
func DeleteUser(username string) error {
	return db.MyDB.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(username))
	})
}
