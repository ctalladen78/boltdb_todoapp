package store

import (
	"encoding/binary"
	"log"
	"time"

	bolt "github.com/etcd-io/bbolt"
)

var taskBucket = []byte("tasks")

type Store struct {
	db *bolt.DB
}

type Task struct {
	Key   int
	Value string
}

func (s *Store) Init(dbPath string) error {
	var err error
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		return err
	}

	s.db = db
	log.Printf("DBPATH  %s", s.db)
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		log.Printf("INIT SUCCESS")
		return err
	})
}

func (s *Store) CreateTask(dbPath string, task string) (int, error) {
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		return 0, err
	}
	log.Printf("XXXXX DB %s", s.db)
	var id int
	// create write transaction
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		id64, _ := bucket.NextSequence()
		id = int(id64)
		key := itob(id)                      // create key
		return bucket.Put(key, []byte(task)) // save to bucket
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Store) AllTasks(dbPath string) ([]Task, error) {
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		return nil, err
	}
	var tasks []Task
	err = db.View(func(tx *bolt.Tx) error {
		buck := tx.Bucket(taskBucket)
		cur := buck.Cursor()
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v), // making sure its immutable
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *Store) DeleteTask(dbPath string, key int) error {
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		buck := tx.Bucket(taskBucket)
		return buck.Delete(itob(key))
	})
	return err
}

// int to byte slice
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// byte slice to int
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
