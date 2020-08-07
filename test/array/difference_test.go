package main

import (
	"reflect"
	"testing"

	"github.com/ypankaj007/go-util/array"
)

func TestDifference(t *testing.T) {
	arr1 := []interface{}{8, 11, 15}
	arr2 := []interface{}{4, 6, 8, 0, 2, 32}
	arr3 := []interface{}{3, 5, 7, 9, 4, 2}
	differenceVal := []interface{}{11, 15}
	differenceRes := array.Difference(arr1, arr2, arr3)
	if !reflect.DeepEqual(differenceRes, differenceVal) {
		t.Errorf("difference was incorrect, got: %v, want: %v.", differenceRes, differenceVal)
	}

	differenceRes = array.Difference(arr2, arr3)
	if !reflect.DeepEqual(differenceRes, []interface{}{6, 8, 0, 32}) {
		t.Errorf("difference was incorrect, got: %v, want: %v.", differenceRes, nil)
	}
	differenceRes = array.Difference(arr2, arr3)
	diff := []interface{}{6, 8, 0, 32}
	if !reflect.DeepEqual(differenceRes, diff) {
		t.Errorf("difference was incorrect, got: %v, want: %v.", differenceRes, diff)
	}

	differenceRes = array.Difference(arr3, arr2)
	diff = []interface{}{3, 5, 7, 9}
	if !reflect.DeepEqual(differenceRes, diff) {
		t.Errorf("difference was incorrect, got: %v, want: %v.", differenceRes, diff)
	}

	differenceRes = array.Difference(arr1, arr1)
	if differenceRes != nil {
		t.Errorf("difference was incorrect, got: %v, want: %v.", differenceRes, nil)
	}
}
