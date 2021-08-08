package utils

import (
	"reflect"
	"testing"
)

func TestReverseMap(t *testing.T) {
	tables := []struct {
		input    map[string]int
		expected map[int]string
	}{
		{map[string]int{"key1": 1, "key2": 2, "key3": 3}, map[int]string{1: "key1", 2: "key2", 3: "key3"}},
		{map[string]int{"a": 1}, map[int]string{1: "a"}},
		{map[string]int{}, map[int]string{}},
		{nil, nil},
	}

	for _, table := range tables {
		t.Logf("input: %v", table.input)
		result := ReverseMap(table.input)
		if !reflect.DeepEqual(table.expected, result) {
			t.Errorf("\tWrong result! Actual is %v but %v expected", result, table.expected)
		}
	}
}
