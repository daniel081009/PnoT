package db_file

import (
	"PnoT/db"
	"PnoT/file"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

func GetFile(username string, path string) (file.File, error) {
	var content []byte

	err := db.MyDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(username))
		if bucket == nil {
			return nil
		}
		content = bucket.Get([]byte(path))
		return nil
	})
	if err != nil {
		return file.File{}, err
	}

	file := file.File{}
	err = file.LoadFile(content)
	if err != nil {
		fmt.Println(err, content)
		return file, err
	}
	return file, err
}

func GetAllFiles(username string) (map[string][]file.File, error) {
	files := make(map[string][]file.File)
	err := db.MyDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(username))
		if bucket == nil {
			return nil
		}
		return bucket.ForEach(func(k, v []byte) error {
			file := file.File{}
			e := file.LoadFile(v)
			if e != nil {
				return e
			}
			files[string(k)] = append(files[string(k)], file)
			return nil
		})
	})
	return files, err
}

func FileExists(username string, path string) bool {
	var exists bool
	db.MyDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(username))
		if bucket == nil {
			return nil
		}
		exists = bucket.Get([]byte(path)) != nil
		return nil
	})
	return exists
}

func AddFile(username string, path string, data string, public bool) error {
	return db.MyDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(username))
		if err != nil {
			return err
		}
		file := file.CreateFile(username, path, data, public)

		fmt.Println(file)
		b, e := file.ToByte()
		if e != nil {
			return e
		}
		return bucket.Put([]byte(file.Path), b)
	})
}
func UpdateFile(username string, path string, data string, public bool) error {
	return db.MyDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(username))
		if err != nil {
			return err
		}
		file := file.File{}
		b := bucket.Get([]byte(path))
		e := file.LoadFile(b)
		if e != nil {
			return e
		}
		file.SaveHistory()
		file.Content = data
		file.Edit_date = time.Now()
		file.Public = public
		ba, e := file.ToByte()
		if e != nil {
			return e
		}
		return bucket.Put([]byte(file.Path), ba)
	})
}
func DeleteFile(username string, key []byte) error {
	return db.MyDB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(username))
		if bucket == nil {
			return nil
		}
		return bucket.Delete(key)
	})
}
