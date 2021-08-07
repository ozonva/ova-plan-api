package utils

import "errors"

// ReverseMap reverse sting->int map to int->string map.
// If input map has one or more same values function calls panic.
// input map must be initialized.
func ReverseMap(input map[string]int) (map[int]string, error) {
	if input == nil {
		return nil, errors.New("uninitialized map passed")
	}

	resultMap := make(map[int]string, len(input))

	for key, value := range input {
		if _, ok := resultMap[value]; ok {
			panic("One or more values in input map have the same values. It's impossible to reverse map.")
		}
		resultMap[value] = key
	}

	return resultMap, nil
}
