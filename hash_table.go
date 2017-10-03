package go_data_structures

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
	"sync"
)

type HashTable struct {
	lock    sync.RWMutex
	storage []interface{}
}

type HashTableItem struct {
	key   interface{}
	value interface{}
	next  *HashTableItem
}

func NewHashTable() *HashTable {
	return &HashTable{sync.RWMutex{}, make([]interface{}, 128)}
}

func (l *HashTable) GetValue(key interface{}) (interface{}, error) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	item, err := l.getHashTableItem(key)

	if err != nil {
		return nil, err
	}

	if item != nil {
		return item.value, nil
	}

	return nil, nil
}

func (l *HashTable) getHashTableItem(key interface{}) (*HashTableItem, error) {
	index, err := l.getHashIndex(key)

	if err != nil {
		return nil, err
	}

	item := l.storage[index].(*HashTableItem)

	for {
		if item.key == key {
			return item, nil
		}

		if item.next == nil {
			break
		}

		item = item.next
	}

	return nil, nil
}

func (l *HashTable) getHashIndex(key interface{}) (uint32, error) {
	code, err := hashCode(key)

	if err != nil {
		return 0, err
	}

	return code / uint32(len(l.storage)), nil
}

func hashCode(key interface{}) (uint32, error) {
	h := fnv.New32a()
	bytes, err := getBytes(key)

	if err != nil {
		return 0, err
	}

	h.Write(bytes)
	return h.Sum32(), nil
}

func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
