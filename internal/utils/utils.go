package utils

import (
	"errors"
)

// ReverseMap reverse sting->int map to int->string map.
// If input map has one or more same values function calls panic.
func ReverseMap(input map[string]int) map[int]string {
	if input == nil {
		return nil
	}

	resultMap := make(map[int]string, len(input))

	for key, value := range input {
		if _, ok := resultMap[value]; ok {
			panic("One or more values in input map have the same values. It's impossible to reverse map.")
		}
		resultMap[value] = key
	}

	return resultMap
}

// SplitSlice splits slice to several batches with size equals batchSize (except last)
func SplitSlice(input []int, batchSize int) ([][]int, error) {
	if input == nil {
		return nil, nil
	}

	if batchSize < 1 {
		return nil, errors.New("batchSize must be greater than 0")
	}

	var batched [][]int

	for i := 0; i < len(input); i += batchSize {
		rightBound := i + batchSize

		if rightBound > len(input) {
			rightBound = len(input)
		}
		batched = append(batched, input[i:rightBound])
	}
	return batched, nil
}

// FilterSlice creates new slice with elements from input without elements from valuesToDelete
func FilterSlice(input []int, valuesToDelete []int) []int {
	if input == nil {
		return nil
	}

	var result []int

	if valuesToDelete == nil || len(valuesToDelete) == 0 {
		copy(result, input)
		return input
	}

	type void struct{}
	var voidValue void

	setToDelete := make(map[int]void)
	for _, val := range valuesToDelete {
		setToDelete[val] = voidValue
	}

	for _, val := range input {
		if _, exists := setToDelete[val]; !exists {
			result = append(result, val)
		}
	}

	return result
}
