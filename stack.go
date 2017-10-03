package go_data_structures

import (
	"errors"
	"sync"
)

type Stack struct {
	lock    sync.Mutex
	length  int
	storage []interface{}
}

func NewStack() *Stack {
	return &Stack{sync.Mutex{}, 0, make([]interface{}, 128)}
}

func (s *Stack) Push(value interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.length == len(s.storage) {
		s.redimStorage()
	}

	s.storage[s.length] = value
	s.length++
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.length == 0 {
		return 0, errors.New("Stack is empty")
	}

	s.length--
	return s.storage[s.length], nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.length == 0 {
		return 0, errors.New("Stack is empty")
	}

	return s.storage[s.length-1], nil
}

func (s *Stack) redimStorage() {
	auxStorage := make([]interface{}, s.length*2)
	copy(s.storage, auxStorage)
	s.storage = auxStorage
}
