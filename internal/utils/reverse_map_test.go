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
		result := ReverseMap(table.input)
		t.Logf("Testing Input: %v", table.input)
		if !reflect.DeepEqual(table.expected, result) {
			t.Errorf("Wrong result! Actual is %v but %v expected", result, table.expected)
		}
	}
}
