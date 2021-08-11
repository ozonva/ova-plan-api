package utils

// FilterSlice creates new slice with elements from input without elements from valuesToDelete
func FilterSlice(input []int, valuesToDelete []int) []int {
	if input == nil {
		return nil
	}
	result := make([]int, 0)
	if valuesToDelete == nil || len(valuesToDelete) == 0 {
		copy(result, input)
		return input
	}

	setToDelete := sliceToSet(valuesToDelete)

	for _, val := range input {
		if _, exists := setToDelete[val]; !exists {
			result = append(result, val)
		}
	}

	return result
}

func sliceToSet(slice []int) map[int]struct{} {
	set := make(map[int]struct{})
	for _, val := range slice {
		set[val] = struct{}{}
	}

	return set
}
