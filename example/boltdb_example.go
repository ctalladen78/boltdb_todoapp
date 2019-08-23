package example

import (
	"time"

	bolt "github.com/etcd-io/bbolt"
	// TODO migrate to storm not bolthold
	// TODO search for asdine/storm code examples
	// "models/data"
)

type model struct {
	KeyId   []byte
	Payload []byte
}

const file = "store.db"

// will be exported
type bboltStore struct {
	db *bolt.DB
}

func (this *bboltStore) Init() error {
	// store := config.Options("store.db")
	db, err := bolt.Open(file, 0666, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		return err
	}

	// setup transaction
	// err = this.db.Update(func(tx *bolt.Tx) error {
	//  _, err := tx.CreateBucketIfNotExists([]byte("bucket_name"))
	//})

	return err
}

func (this *bboltStore) SaveData(mod *model) error {
	err := this.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("bucket_name"))
		return b.Put([]byte(mod.KeyId), mod.Payload)
	})
	return err
}

func (this *bboltStore) GetData() ([]*model, error) {
	//mod := &model{}
	var result []*model
	err := this.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("bucket_name"))
		cursor := bucket.Cursor()
		limit := 16
		idx := 0
		var temp []*model
		for k, v := cursor.First(); k != nil && idx < limit; k, v = cursor.Next() {
			d := &model{
				KeyId:   k,
				Payload: v,
			}
			temp = append(temp, d)
			err := cursor.Delete() // TODO research this
			if err != nil {
				return err
			}
			idx++
		}
		result = temp
		return nil
	})
	return result, err
}

func QueryDataByHash(hash []byte) error { return nil }
