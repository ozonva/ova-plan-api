package utils

import (
	"errors"
	"reflect"
	"testing"
)

func TestSplitSlice(t *testing.T) {
	tables := []struct {
		inputSlice     []int
		inputBatchSize int
		expected       [][]int
		expectedErr    error
	}{
		{[]int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {3, 4}}, nil},
		{[]int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}, nil},
		{[]int{1, 2, 3}, 3, [][]int{{1, 2, 3}}, nil},
		{[]int{1, 2, 3}, 0, nil, errors.New("batchSize must be greater than 0")},
		{[]int{1, 2}, 3, [][]int{{1, 2}}, nil},
		{[]int{1, 2, 3}, 2, [][]int{{1, 2}, {3}}, nil},
		{[]int{}, 2, [][]int{}, nil},
		{nil, 1, nil, nil},
	}

	for _, table := range tables {
		t.Logf("input: %v %v", table.inputSlice, table.inputBatchSize)
		result, err := SplitSlice(table.inputSlice, table.inputBatchSize)

		if !reflect.DeepEqual(table.expectedErr, err) {
			t.Errorf("\tWrong error! Actual is \"%v\" but \"%v\" expected", err, table.expectedErr)
		}
		if !reflect.DeepEqual(table.expected, result) {
			t.Errorf("\tWrong result! Actual is %v but %v expected", result, table.expected)
		}
	}
}
