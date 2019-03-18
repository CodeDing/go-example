package main

import (
	"errors"
	"fmt"
	"github.com/dgraph-io/badger"
)

type BadgerDB struct {
	db *badger.DB
}

func NewBadgerDB(dir string) *BadgerDB {
	opts := badger.DefaultOptions
	opts.Dir = dir
	opts.ValueDir = dir
	db, err := badger.Open(opts)
	if err != nil {
		fmt.Printf("New badgerrDB failed!\n")
		return nil
	}
	return &BadgerDB{db: db}
}

func (badgerDB *BadgerDB) Get(key []byte) ([]byte, error) {
	if badgerDB.db == nil {
		return nil, errors.New("badger db not open!")
	}
	var val []byte
	err := badgerDB.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		err = item.Value(func(v []byte) error {
			val = append([]byte{}, v...)
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return val, nil

}

//返回第一个命中的prefix key
func (badgerDB *BadgerDB) GetPrefix(prefix []byte) ([]byte, error) {
	if badgerDB.db == nil {
		return nil, errors.New("badger db not open!")
	}
	var val []byte
	err := badgerDB.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		//for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
		it.Seek(prefix)
		if it.ValidForPrefix(prefix) {
			item := it.Item()
			err := item.Value(func(v []byte) error {
				val = append([]byte{}, v...)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (badgerDB *BadgerDB) Set(key, value []byte) error {
	if badgerDB.db == nil {
		return errors.New("badger db not open!")
	}
	err := badgerDB.db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, value)
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

func (badgerDB *BadgerDB) BulkUpdate(kv map[string]string) error {
	if badgerDB.db == nil {
		return errors.New("badger db not open!")
	}

	txn := badgerDB.db.NewTransaction(true)
	defer txn.Discard()
	var err error
	for k, v := range kv {
		if err = txn.Set([]byte(k), []byte(v)); err == badger.ErrTxnTooBig {
			fmt.Printf("badger txn too big\n")
			err = txn.Commit()
			if err != nil {
				//TODO: debug log
				continue
			}
			txn = badgerDB.db.NewTransaction(true)
			err = txn.Set([]byte(k), []byte(v))
			if err != nil {
				//TODO: debug log
			}
		}
	}
	err = txn.Commit()
	if err != nil {
		//TODO: debug log
	}
	return err

}

func (badgerDB *BadgerDB) Close() {
	if badgerDB.db == nil {
		return
	}
	badgerDB.db.Close()
	badgerDB.db = nil
	return
}

func main() {

	fmt.Println("================    start   ==================")
	badgerDB := NewBadgerDB("/tmp/badgerdb")
	defer badgerDB.Close()
	keys := []string{"key1", "key2", "key3"}
	updates := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	for k, v := range updates {
		err := badgerDB.Set([]byte(k), []byte(v))
		fmt.Printf("badgerDB set failed, key: %s, value:%s, err: %v\n", k, v, err)
	}
	_ = badgerDB.Set([]byte("key1"), []byte("value1_new"))
	for _, key := range keys {
		val, err := badgerDB.Get([]byte(key))
		if err == nil {
			fmt.Printf("key:%s, value: %s\n", key, string(val))
		}
	}

	val, err := badgerDB.GetPrefix([]byte("key"))
	if err == nil {
		fmt.Printf("badgerDB get prefix: key, value:%s\n", string(val))
	}

	const N = 1000000
	bulk_keys := make([]string, 0, N)
	bulk_updates := make(map[string]string, N)
	for i := 0; i < N; i++ {
		bulk_keys = append(bulk_keys, fmt.Sprintf("key%d", i))
		bulk_updates[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
	}
	if err := badgerDB.BulkUpdate(bulk_updates); err != nil {
		fmt.Printf("badger bulk updates failed, err:%v\n", err)
	}

	LAST := 10
	for i := N - LAST; i < N; i++ {
		key := fmt.Sprintf("key%d", i)
		val, err := badgerDB.Get([]byte(key))
		if err == nil {
			fmt.Printf("key:%s, value: %s\n", key, string(val))
		}

	}
	fmt.Println("================    end   ==================")
}
