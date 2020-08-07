package main

import (
	"testing"

	"github.com/ypankaj007/go-util/array"
)

func TestIndexOf(t *testing.T) {
	arr := []interface{}{4, 6, 8, 0, 2, 32, 6}
	idx := array.IndexOf(arr, 0, func(v interface{}) bool {
		return v.(int) == 8
	})
	if idx != 2 {
		t.Errorf("IndexOf was incorrect, got: %d, want: %d.", idx, 2)
	}
	idx = array.IndexOf(arr, 0, func(v interface{}) bool {
		return v.(int) == 2
	})
	if idx != 4 {
		t.Errorf("IndexOf was incorrect, got: %d, want: %d.", idx, 4)
	}
	idx = array.IndexOf(arr, 0, func(v interface{}) bool {
		return v.(int) == 10
	})
	if idx != -1 {
		t.Errorf("IndexOf was incorrect, got: %d, want: %d.", idx, -1)
	}
	idx = array.IndexOf(arr, 0, func(v interface{}) bool {
		return v.(int) == 6
	})
	if idx != 1 {
		t.Errorf("IndexOf was incorrect, got: %d, want: %d.", idx, 1)
	}
	idx = array.IndexOf(arr, 2, func(v interface{}) bool {
		return v.(int) == 6
	})
	if idx != 6 {
		t.Errorf("IndexOf was incorrect, got: %d, want: %d.", idx, 6)
	}
	idx = array.IndexOf(nil, 2, func(v interface{}) bool {
		return v.(int) == 6
	})
	if idx != -1 {
		t.Errorf("IndexOf was incorrect, got: %d, want: %d.", idx, -1)
	}
}
