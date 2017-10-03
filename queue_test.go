package go_data_structures

import (
	"math/rand"
	"testing"
)

func TestEnqueueInteger(t *testing.T) {
	q := NewQueue()
	err := q.Enqueue(1)

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	if q.length != 1 {
		t.Errorf("Test failed, expect length to be 1, got: '%s'", q.length)
	}
}

func TestEnqueueString(t *testing.T) {
	q := NewQueue()
	err := q.Enqueue("Test")

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	if q.length != 1 {
		t.Errorf("Test failed, expect length to be 1, got: '%s'", q.length)
	}
}

func TestEnqueueFullQueueStorage(t *testing.T) {
	q := NewQueue()
	expected := len(q.storage) + 1

	for i := 0; i < len(q.storage); i++ {
		q.Enqueue(i)
	}

	q.Enqueue("test")

	if q.length != expected {
		t.Errorf("Test failed, expect length to be '%s', got: '%s'", expected, q.length)
	}
}

func TestEnqueueFullQueueStorageAfterDequeues(t *testing.T) {
	q := NewQueue()
	expected := len(q.storage) + 1

	for i := 0; i < len(q.storage); i++ {
		q.Enqueue(i)
	}

	for i := 0; i < len(q.storage)/2; i++ {
		q.Dequeue()
		q.Enqueue(i)
	}

	q.Enqueue("test")
	value, _ := q.Peek()

	if q.length != expected {
		t.Errorf("Test failed, expect length to be '%s', got: '%s'", expected, q.length)
	}

	if value != "test" {
		t.Errorf("Test failed, expect value to be '%s', got: '%s'", "test", value)
	}
}

func TestDequeueBackDoFirstIndex(t *testing.T) {
	q := NewQueue()
	expected := len(q.storage) + 1

	for i := 0; i < len(q.storage); i++ {
		q.Enqueue(i)
	}

	for i := 0; i < len(q.storage); i++ {
		q.Dequeue()
		q.Enqueue(i)
	}

	q.Enqueue("test")
	value, _ := q.Peek()

	if q.length != expected {
		t.Errorf("Test failed, expect length to be '%s', got: '%s'", expected, q.length)
	}

	if value != "test" {
		t.Errorf("Test failed, expect value to be '%s', got: '%s'", "test", value)
	}
}

func TestDequeueEmptyQueue(t *testing.T) {
	q := NewQueue()
	value, err := q.Dequeue()

	expected := 0

	if err == nil {
		t.Errorf("Test failed, error must not be nil")
	}

	if value != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, value)
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue()
	expected := "Test"

	q.Enqueue(expected)
	value, err := q.Dequeue()

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	if value != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, value)
	}

	if q.length != 0 {
		t.Errorf("Test failed, stack length must be 0")
	}
}

func TestPeekEmptyQueue(t *testing.T) {
	q := NewQueue()
	value, err := q.Peek()

	expected := 0

	if err == nil {
		t.Errorf("Test failed, error must not be nil")
	}

	if value != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, value)
	}
}

func TestQueuePeek(t *testing.T) {
	q := NewQueue()
	expected := "Test"

	q.Enqueue(expected)
	value, err := q.Peek()

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	if value != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, value)
	}
}

func TestFIFO(t *testing.T) {
	q := NewQueue()
	value1 := 1
	value2 := "test"
	value3 := true
	value4 := [3]int{1, 2, 3}

	q.Enqueue(value1)
	q.Enqueue(value2)
	q.Enqueue(value3)
	q.Enqueue(value4)

	actual1, _ := q.Dequeue()

	if value4 != actual1 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", value4, actual1)
	}

	actual2, _ := q.Dequeue()

	if value3 != actual2 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", value3, actual2)
	}

	actual3, _ := q.Dequeue()

	if value2 != actual3 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", value2, actual3)
	}

	actual4, _ := q.Dequeue()

	if value1 != actual4 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", value1, actual4)
	}
}

func TestEnqueueDequeueThreadSafety(t *testing.T) {
	q := NewQueue()
	c := make(chan bool)
	gr := 50

	for i := 0; i < gr; i++ {
		go EnqueueRoutine(q, c)
	}

	for i := 0; i < gr; i++ {
		<-c
	}

	expected := gr * 1000

	if q.length != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, q.length)
	}
}

func EnqueueRoutine(q *Queue, c chan bool) {
	for i := 0; i < 1000; i++ {
		q.Enqueue(rand.Int())
	}

	c <- true
}
