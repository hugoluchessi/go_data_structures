package go_data_structures

import (
	"strconv"
	"testing"
)

func TestAddFirstInteger(t *testing.T) {
	var err error
	l := NewLinkedList()
	length := 10

	for i := 0; i < length; i++ {
		err = l.AddFirst(i)
	}

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	if l.length != length {
		t.Errorf("Test failed, expect length to be '%s', got: '%s'", length, l.length)
	}
}

func TestAddFirstString(t *testing.T) {
	var err error
	l := NewLinkedList()
	length := 10

	for i := 0; i < length; i++ {
		err = l.AddFirst(strconv.Itoa(i))
	}

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	if l.length != length {
		t.Errorf("Test failed, expect length to be '%s', got: '%s'", length, l.length)
	}
}

func TestRemoveInteger(t *testing.T) {
	var err error
	l := NewLinkedList()
	length := 10

	for i := 0; i < length; i++ {
		err = l.AddFirst(i)
	}

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	if l.length != length {
		t.Errorf("Test failed, expect length to be '%s', got: '%s'", length, l.length)
	}

	for i := 0; i < length; i++ {
		err = l.Remove(i)
	}

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	if l.length != 0 {
		t.Errorf("Test failed, expect length to be '%s', got: '%s'", 0, l.length)
	}
}

func TestRemoveString(t *testing.T) {
	var err error
	l := NewLinkedList()
	length := 10

	for i := 0; i < length; i++ {
		err = l.AddFirst(strconv.Itoa(i))
	}

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	if l.length != length {
		t.Errorf("Test failed, expect length to be '%s', got: '%s'", length, l.length)
	}

	for i := 0; i < length; i++ {
		err = l.Remove(strconv.Itoa(i))
	}

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	if l.length != 0 {
		t.Errorf("Test failed, expect length to be '%s', got: '%s'", 0, l.length)
	}
}

func TestHasValueInteger(t *testing.T) {
	l := NewLinkedList()
	value := 5
	length := 10

	for i := 0; i < length; i++ {
		l.AddFirst(i)
	}

	if !l.HasValue(value) {
		t.Errorf("Test failed, added value not found")
	}

	l.Remove(value)

	if l.HasValue(value) {
		t.Errorf("Test failed, removed value found")
	}
}
