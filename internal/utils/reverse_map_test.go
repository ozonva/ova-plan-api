package utils

import (
	"reflect"
	"strings"
	"testing"
)

func TestReverseMap(t *testing.T) {
	tables := []struct {
		input           map[string]int
		expected        map[int]string
		errMustContains string
	}{
		{map[string]int{"key1": 1, "key2": 2, "key3": 3}, map[int]string{1: "key1", 2: "key2", 3: "key3"}, ""},
		{map[string]int{"a": 1}, map[int]string{1: "a"}, ""},
		{map[string]int{}, map[int]string{}, ""},
		{nil, nil, ""},
		{map[string]int{"a": 2, "b": 2, "c": 3}, nil, "one or more values in input map have the same values"},
	}

	for _, table := range tables {
		t.Logf("input: %v", table.input)
		result, err := ReverseMap(table.input)

		if !errorContains(err, table.errMustContains) {
			t.Errorf("\tExpected error contains \"%v\" but got \"%v\"", table.errMustContains, err)
		}
		if !reflect.DeepEqual(table.expected, result) {
			t.Errorf("\tWrong result! Actual is %v but %v expected", result, table.expected)
		}
	}
}

func errorContains(err error, expected string) bool {
	if err == nil {
		return expected == ""
	}
	if expected == "" {
		return false
	}
	return strings.Contains(err.Error(), expected)
}
