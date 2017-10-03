package go_data_structures

import (
	"sync"
)

type LinkedList struct {
	lock   sync.RWMutex
	head   *LinkedListItem
	length int
}

type LinkedListItem struct {
	next  *LinkedListItem
	value interface{}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{sync.RWMutex{}, &LinkedListItem{nil, nil}, 0}
}

func (l *LinkedList) AddFirst(value interface{}) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	item := &LinkedListItem{nil, value}

	item.next = l.head.next
	l.head.next = item
	l.length++

	return nil
}

func (l *LinkedList) HasValue(value interface{}) bool {
	l.lock.RLock()
	defer l.lock.RUnlock()

	item := l.head

	for item.next != nil {
		if item.next.value == value {
			return true
		}

		item = item.next
	}

	return false
}

func (l *LinkedList) Remove(value interface{}) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	item := l.head

	for item.next != nil {
		if item.next.value == value {
			item.next = item.next.next
			l.length--
			break
		}

		item = item.next
	}

	return nil
}
