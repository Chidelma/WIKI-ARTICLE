package main

import "testing"

func TestValueExist(t *testing.T) {

	setValue("rest_api", "REpresentational State Transfer Application Programming Interface")

	exists := valueExist("rest_api")

	if !exists {
		t.Errorf("Value should exist, got: %t, want: %t.", exists, true)
	}
}

func TestValueNotExist(t *testing.T) {

	exists := valueExist("wiki")

	if exists {
		t.Errorf("Value should not exist, got: %t, want: %t.", exists, false)
	}
}

func TestInitDataLength(t *testing.T) {

	keys := retriveAllKeys()

	if len(keys) == 0 {
		t.Errorf("Length must be 1, got: %d, want: %d.", len(keys), 1)
	} 
}

func TestDataLengthAfterInsert(t *testing.T) {

	prevDataLen := len(retriveAllKeys())

	setValue("wiki", "A wiki is a knowledge base website")

	nextDataLen := len(retriveAllKeys())

	if !(nextDataLen > prevDataLen) {
		t.Errorf("Length must be 2, got: %d, want: %d.", nextDataLen, 2)
	}
}

func TestDataLengthAfterSameKeyInsert(t *testing.T) {

	prevDataLen := len(retriveAllKeys())

	setValue("rest_api", "A wiki is a knowledge base website")

	nextDataLen := len(retriveAllKeys())

	if !(nextDataLen == prevDataLen) {
		t.Errorf("Length must be 2, got: %d, want: %d.", nextDataLen, 2)
	}
}

func TestGetValue(t *testing.T) {

	value := getValue("rest_api")

	if value == "" {
		t.Errorf("Value should exist, got: %s, want: %s.", value, "A wiki is a knowledge base website")
	}
}