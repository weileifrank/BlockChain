package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("bl"))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("bl"))
			if err != nil {
				panic(err)
			}
		}
		bucket.Put([]byte("name"),[]byte("frank"))
		bucket.Put([]byte("age"),[]byte("18"))
		return nil
	})
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("bl"))
		if bucket == nil {
			panic(err)
		}
		name := bucket.Get([]byte("name"))
		age := bucket.Get([]byte("age"))
		fmt.Printf("name=%s       age=%s",name,age)
		return nil
	})
}
