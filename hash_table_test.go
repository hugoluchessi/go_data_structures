package go_data_structures

import (
	"testing"
)

func TestSetValueIntegerKeyIntegerValue(t *testing.T) {
	var err error
	var actualValue interface{}

	key := 1
	value := 10

	h := NewHashTable()

	err = h.SetValue(key, value)

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	actualValue, err = h.GetValue(key)

	if actualValue.(int) != value {
		t.Errorf("Test failed, expect value to be '%s', got: '%s'", value, actualValue)
	}
}

func TestSetValueStringKeyIntegerValue(t *testing.T) {
	var err error
	var actualValue interface{}

	key := "string key"
	value := 10

	h := NewHashTable()

	err = h.SetValue(key, value)

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	actualValue, err = h.GetValue(key)

	if actualValue.(int) != value {
		t.Errorf("Test failed, expect value to be '%s', got: '%s'", value, actualValue)
	}
}

func TestSetValueStringKeyStringValue(t *testing.T) {
	var err error
	var actualValue interface{}

	key := "string key"
	value := "string value"

	h := NewHashTable()

	err = h.SetValue(key, value)

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	actualValue, err = h.GetValue(key)

	if actualValue.(string) != value {
		t.Errorf("Test failed, expect value to be '%s', got: '%s'", value, actualValue)
	}
}

func TestSetExistingValueStringKeyStringValue(t *testing.T) {
	var err error
	var actualValue interface{}

	key := "string key"
	value := "string value"
	value2 := "value string"

	h := NewHashTable()

	err = h.SetValue(key, value)

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	actualValue, err = h.GetValue(key)

	if actualValue.(string) != value {
		t.Errorf("Test failed, expect value to be '%s', got: '%s'", value, actualValue)
	}

	err = h.SetValue(key, value2)
	actualValue, err = h.GetValue(key)

	if actualValue.(string) != value2 {
		t.Errorf("Test failed, expect value to be '%s', got: '%s'", value2, actualValue)
	}
}

func TestSetMoreValuesThanBuckets(t *testing.T) {
	h := NewHashTable()

	for i := 0; i < 5000; i++ {
		_ = h.SetValue(i, i)
	}

	for i := 0; i < 5000; i++ {
		value, _ := h.GetValue(i)

		if value.(int) != i {
			t.Errorf("Test failed, expect value to be '%s', got: '%s'", i, value)
		}
	}
}

func TestGetInexistingValues(t *testing.T) {
	h := NewHashTable()

	value, err := h.GetValue("anything")

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	if value != nil {
		t.Errorf("Test failed, expect value to be '%s', got: '%s'", nil, value)
	}
}

func TestRemoveKey(t *testing.T) {
	var err error
	var actualValue interface{}

	key := "string key"
	value := "string value"

	h := NewHashTable()

	err = h.SetValue(key, value)

	if err != nil {
		t.Errorf("Test failed, error must be nil")
	}

	actualValue, err = h.GetValue(key)

	if actualValue.(string) != value {
		t.Errorf("Test failed, expect value to be '%s', got: '%s'", value, actualValue)
	}

	err = h.RemoveKey(key)
	actualValue, err = h.GetValue(key)

	if actualValue != nil {
		t.Errorf("Test failed, expect value to be '%s', got: '%s'", nil, actualValue)
	}
}

func TestRemoveALotOfKeys(t *testing.T) {
	h := NewHashTable()

	for i := 0; i < 500; i++ {
		_ = h.SetValue(i, i)
	}

	for i := 0; i < 500; i++ {
		err := h.RemoveKey(i)

		if err != nil {
			t.Errorf("Test failed, expect value to be '%s', got: '%s'", nil, err)
		}
	}

	for i := 0; i < 500; i++ {
		_ = h.SetValue(i, i)
	}

	for i := 500; i >= 0; i-- {
		err := h.RemoveKey(i)

		if err != nil {
			t.Errorf("Test failed, expect value to be '%s', got: '%s'", nil, err)
		}
	}
}
