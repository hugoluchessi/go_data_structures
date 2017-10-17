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
	length  int
}

type HashTableItem struct {
	key   interface{}
	value interface{}
	next  *HashTableItem
}

func NewHashTable() *HashTable {
	return &HashTable{sync.RWMutex{}, make([]interface{}, 64), 0}
}

func (h *HashTable) GetValue(key interface{}) (interface{}, error) {
	h.lock.RLock()
	defer h.lock.RUnlock()

	item, err := h.getHashTableItem(key)

	if err != nil {
		return nil, err
	}

	if item != nil {
		return item.value, nil
	}

	return nil, nil
}

func (h *HashTable) SetValue(key interface{}, value interface{}) error {
	h.lock.Lock()
	defer h.lock.Unlock()

	err := h.setValueUnsafe(key, value)

	if err != nil {
		return err
	}

	h.length++

	if h.length > len(h.storage)*2 {
		err = h.redistributeValues()

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *HashTable) RemoveKey(key interface{}) error {
	h.lock.Lock()
	defer h.lock.Unlock()

	index, err := h.getHashIndex(key)

	if err != nil {
		return err
	}

	if h.storage[index] == nil {
		return nil
	}

	item := h.storage[index].(*HashTableItem)

	if item.key == key {
		if item.next != nil {
			h.storage[index] = item.next
		} else {
			h.storage[index] = nil
		}

		return nil
	}

	for {
		previousItem := item
		item = item.next

		if item.key == key {
			if item.next != nil {
				previousItem.next = item.next
			} else {
				previousItem.next = nil
			}

			return nil
		}

		if item.next == nil {
			return nil
		}
	}
}

func (h *HashTable) setValueUnsafe(key interface{}, value interface{}) error {
	item, err := h.getHashTableItem(key)

	if err != nil {
		return err
	}

	if item != nil {
		item.value = value
		return nil
	}

	index, err := h.getHashIndex(key)

	if err != nil {
		return err
	}

	if h.storage[index] != nil {
		item = h.storage[index].(*HashTableItem)
	}

	h.storage[index] = &HashTableItem{key, value, item}

	return nil
}

func (h *HashTable) getHashTableItem(key interface{}) (*HashTableItem, error) {
	var item *HashTableItem
	index, err := h.getHashIndex(key)

	if err != nil {
		return nil, err
	}

	if h.storage[index] == nil {
		return nil, nil
	}

	item = h.storage[index].(*HashTableItem)

	for {
		if item.key == key {
			return item, nil
		}

		if item.next == nil {
			return nil, nil
		}

		item = item.next
	}
}

func (h *HashTable) getHashIndex(key interface{}) (int, error) {
	code, err := hashCode(key)

	if err != nil {
		return 0, err
	}

	return int(code) % len(h.storage), nil
}

func (h *HashTable) redistributeValues() error {
	currentLength := len(h.storage)
	currentStorage := make([]interface{}, currentLength)
	copy(currentStorage, h.storage)

	h.storage = make([]interface{}, h.length*2)

	for i := 0; i < currentLength; i++ {
		if currentStorage[i] == nil {
			continue
		}

		item := currentStorage[i].(*HashTableItem)

		for item != nil {
			err := h.setValueUnsafe(item.key, item.value)

			if err != nil {
				return err
			}

			item = item.next
		}
	}

	return nil
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
