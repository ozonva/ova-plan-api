package utils

import "errors"

// ReverseMap reverse sting->int map to int->string map.
// If input map has one or more same values function calls panic.
func ReverseMap(input map[string]int) (map[int]string, error) {
	if input == nil {
		return nil, nil
	}

	resultMap := make(map[int]string, len(input))

	for key, value := range input {
		if _, ok := resultMap[value]; ok {
			return nil, errors.New("one or more values in input map have the same values. It's impossible to reverse map")
		}
		resultMap[value] = key
	}

	return resultMap, nil
}
