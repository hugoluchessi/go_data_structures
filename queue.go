package go_data_structures

import (
	"errors"
	"sync"
)

type Queue struct {
	lock    sync.Mutex
	i       int
	e       int
	length  int
	storage []interface{}
}

func NewQueue() *Queue {
	return &Queue{sync.Mutex{}, 0, 0, 0, make([]interface{}, 128)}
}

func (q *Queue) Enqueue(value interface{}) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.length == len(q.storage) {
		q.redimStorage()
	}

	q.storage[q.e] = value
	q.length++
	q.e++

	if q.e >= len(q.storage) {
		q.e = 0
	}

	return nil
}

func (q *Queue) Dequeue() (interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.length == 0 {
		return 0, errors.New("Queue is empty")
	}

	item := q.storage[q.i]

	q.length--
	q.i++

	if q.i >= len(q.storage) {
		q.i = 0
	}

	return item, nil
}

func (q *Queue) Peek() (interface{}, error) {
	if q.length == 0 {
		return 0, errors.New("Queue is empty")
	}

	return q.storage[q.i], nil
}

func (q *Queue) redimStorage() {
	auxStorage := make([]interface{}, q.length*2)
	priorLength := q.length
	auxIndex := 0

	for auxI := q.i; auxIndex < q.length; auxI++ {
		i := auxI

		if auxI >= len(q.storage) {
			i -= len(q.storage)
		}

		auxStorage[auxIndex] = q.storage[i]
		auxIndex++
	}

	q.i = 0
	q.e = priorLength
	q.storage = auxStorage
}
