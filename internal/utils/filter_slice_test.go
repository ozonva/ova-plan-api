package utils

import (
	"reflect"
	"testing"
)

func TestFilterSlice(t *testing.T) {
	tables := []struct {
		inputSlice    []int
		inputToDelete []int
		expected      []int
	}{
		{[]int{1, 2}, []int{1}, []int{2}},
		{[]int{1}, []int{1}, []int{}},
		{[]int{}, []int{}, []int{}},
		{[]int{1, 2}, []int{}, []int{1, 2}},
		{nil, []int{}, nil},
		{[]int{1}, nil, []int{1}},
	}

	for _, table := range tables {
		result := FilterSlice(table.inputSlice, table.inputToDelete)
		t.Logf("input: %v %v", table.inputSlice, table.inputToDelete)
		if !reflect.DeepEqual(table.expected, result) {
			t.Errorf("\tWrong result! Actual is %v but %v expected", result, table.expected)
		}
		if result != nil && table.inputSlice != nil && &result == &table.inputSlice {
			t.Error("New slice must be created, not reused")
		}
	}
}
