package backend

import (
	"github.com/dgraph-io/badger"
	log "github.com/sirupsen/logrus"
)

type Badger struct {
	client *badger.DB
}

func NewBadger(bo badger.Options) *Badger {
	b := &Badger{}
	db, err := badger.Open(bo)
	if err != nil {
		log.Fatal(err)
	}
	b.client = db
	return b
}

func (b Badger) Get(key string) (string, error) {
	var val string
	err := b.client.View(func(txn *badger.Txn) error {
		i, err := txn.Get([]byte(key))
		if err != nil {
			return err
		} else {
			err = i.Value(func(v []byte) error {
				val = string(v)
				return nil
			})
		}
		return nil
	})
	return val, err
}

func (b Badger) GetAll() (map[string]string, error) {
	m := map[string]string{}
	err := b.client.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte("")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := string(item.Key())
			err := item.Value(func(v []byte) error {
				m[k] = string(v)
				return nil
			})
			if err != nil {
				log.Warn(err)
			}
		}
		return nil
	})
	return m, err
}

func (b Badger) Set(key string, value string, ttl int) error {
	err := b.client.Update(func(txn *badger.Txn) error {
		// todo hadle TTL
		e := badger.NewEntry([]byte(key), []byte(value)) // .WithTTL(time.Hour)
		err := txn.SetEntry(e)
		return err
	})
	return err
}

func (b Badger) Close() {
	err := b.client.Close()
	if err != nil {
		log.Error(err)
	}
}