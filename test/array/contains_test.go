package main

import (
	"testing"

	"github.com/ypankaj007/go-util/array"
)

func TestContains(t *testing.T) {
	arr := []interface{}{4, 6, 8, 0, 2, 32, 6}
	result := array.Contains(arr, 0, func(v interface{}) bool {
		return v.(int) == 8
	})
	if result != true {
		t.Errorf("Contains was incorrect, got: %t, want: %t.", result, true)
	}
	result = array.Contains(arr, 0, func(v interface{}) bool {
		return v.(int) == 2
	})
	if result != true {
		t.Errorf("Contains was incorrect, got: %t, want: %t.", result, true)
	}
	result = array.Contains(arr, 0, func(v interface{}) bool {
		return v.(int) == 10
	})
	if result != false {
		t.Errorf("Contains was incorrect, got: %t, want: %t.", result, false)
	}
	result = array.Contains(arr, 0, func(v interface{}) bool {
		return v.(int) == 6
	})
	if result != true {
		t.Errorf("Contains was incorrect, got: %t, want: %t.", result, true)
	}
	result = array.Contains(arr, 2, func(v interface{}) bool {
		return v.(int) == 6
	})
	if result != true {
		t.Errorf("Contains was incorrect, got: %t, want: %t.", result, true)
	}
	result = array.Contains(nil, 2, func(v interface{}) bool {
		return v.(int) == 6
	})
	if result != false {
		t.Errorf("Contains was incorrect, got: %t, want: %t.", result, false)
	}
}
