package main

import (
	"reflect"
	"testing"

	"github.com/ypankaj007/go-util/array"
)

func TestUniq(t *testing.T) {
	arr := []interface{}{4, 6, 8, 0, 2, 32, 6}
	uniqVal := []interface{}{4, 6, 8, 0, 2, 32}
	uniqRes := array.Uniq(arr)
	if !reflect.DeepEqual(uniqRes, uniqVal) {
		t.Errorf("Uniq was incorrect, got: %v, want: %v.", uniqRes, uniqVal)
	}

	uniqRes = array.Uniq(nil)
	if uniqRes != nil {
		t.Errorf("Uniq was incorrect, got: %v, want: %v.", uniqRes, nil)
	}
}
