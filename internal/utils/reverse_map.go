package utils

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
