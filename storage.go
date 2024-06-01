package main
import (
    "fmt"
    "log"
	bolt "go.etcd.io/bbolt"
)

func retrieve(db *bolt.DB, bucket string, key string) string {
    var res []byte
    db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(bucket))
        v := b.Get([]byte(key))
        res = make([]byte, len(v))
        _ = copy(res, v)
        return nil
    })
    return string(res)
}

type Pair struct {
    First []byte
    Second []byte
}

func retrieveAll(db *bolt.DB, bucket string) []Pair {
    var res []Pair
    db.View(func(tx *bolt.Tx) error {
        // Assume bucket exists and has keys
        b := tx.Bucket([]byte(bucket))
        c := b.Cursor()
        for k, v := c.First(); k != nil; k, v = c.Next() {
            //fmt.Printf("key=%s, value=%s\n", k, v)
            res = append(res, Pair { k, v })
        }
        return nil
    })
    return res
}

func update(db *bolt.DB, bucket string, key []byte, value []byte) {
    db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(bucket))
        err := b.Put(key, value)
        return err
    })
}

func createBucket(db *bolt.DB, bucket string) {
    db.Update(func(tx *bolt.Tx) error {
        _, err := tx.CreateBucket([]byte(bucket))
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }
        return nil
    })
}

func initDB() *bolt.DB {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
    createBucket(db, "VM")
    return db
}
