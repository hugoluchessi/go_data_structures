package go_data_structures

import (
	"math/rand"
	"testing"
)

func TestPushInteger(t *testing.T) {
	s := NewStack()
	err := s.Push(1)

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	if s.length != 1 {
		t.Errorf("Test failed, expect length to be 1, got: '%s'", s.length)
	}
}

func TestPushString(t *testing.T) {
	s := NewStack()
	err := s.Push("Test")

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	if s.length != 1 {
		t.Errorf("Test failed, expect length to be 1, got: '%s'", s.length)
	}
}

func TestPushFullStackStorage(t *testing.T) {
	s := NewStack()
	expected := len(s.storage) + 1

	for i := 0; i < len(s.storage); i++ {
		s.Push(i)
	}

	s.Push("test")

	if s.length != expected {
		t.Errorf("Test failed, expect length to be '%s', got: '%s'", expected, s.length)
	}
}

func TestPopEmptyStack(t *testing.T) {
	s := NewStack()
	value, err := s.Pop()

	expected := 0

	if err == nil {
		t.Errorf("Test failed, error must not be nil")
	}

	if value != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, value)
	}
}

func TestPop(t *testing.T) {
	s := NewStack()
	expected := "Test"

	s.Push(expected)
	value, err := s.Pop()

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	if value != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, value)
	}

	if s.length != 0 {
		t.Errorf("Test failed, stack length must be 0")
	}
}

func TestPeekEmptyStack(t *testing.T) {
	s := NewStack()
	value, err := s.Peek()

	expected := 0

	if err == nil {
		t.Errorf("Test failed, error must not be nil")
	}

	if value != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, value)
	}
}

func TestStackPeek(t *testing.T) {
	s := NewStack()
	expected := "Test"

	s.Push(expected)
	value, err := s.Peek()

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	if value != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, value)
	}
}

func TestLIFO(t *testing.T) {
	s := NewStack()
	value1 := 1
	value2 := "test"
	value3 := true
	value4 := [3]int{1, 2, 3}

	s.Push(value1)
	s.Push(value2)
	s.Push(value3)
	s.Push(value4)

	actual1, _ := s.Pop()

	if value4 != actual1 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", value4, actual1)
	}

	actual2, _ := s.Pop()

	if value3 != actual2 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", value3, actual2)
	}

	actual3, _ := s.Pop()

	if value2 != actual3 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", value2, actual3)
	}

	actual4, _ := s.Pop()

	if value1 != actual4 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", value1, actual4)
	}
}

func TestPushPopThreadSafety(t *testing.T) {
	s := NewStack()
	c := make(chan bool)
	gr := 50

	for i := 0; i < gr; i++ {
		go PushRoutine(s, c)
	}

	for i := 0; i < gr; i++ {
		<-c
	}

	expected := gr * 1000

	if s.length != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, s.length)
	}
}

func PushRoutine(s *Stack, c chan bool) {
	for i := 0; i < 1000; i++ {
		s.Push(rand.Int())
	}

	c <- true
}
