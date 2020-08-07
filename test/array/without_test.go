package main

import (
	"reflect"
	"testing"

	"github.com/ypankaj007/go-util/array"
)

func TestWithout(t *testing.T) {
	arr1 := []interface{}{4, 6, 8, 0, 2, 32, 7}
	arr2 := []interface{}{4, 6, 8, 0, 2, 32}
	withoutVal := []interface{}{7}
	withoutRes := array.Without(arr1, arr2)
	if !reflect.DeepEqual(withoutRes, withoutVal) {
		t.Errorf("without was incorrect, got: %v, want: %v.", withoutRes, withoutVal)
	}

	withoutRes = array.Without(nil, nil)
	if withoutRes != nil {
		t.Errorf("without was incorrect, got: %v, want: %v.", withoutRes, nil)
	}

	withoutRes = array.Without(arr1, arr1)
	if withoutRes != nil {
		t.Errorf("without was incorrect, got: %v, want: %v.", withoutRes, nil)
	}
}
