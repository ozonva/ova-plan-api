package utils

import "errors"

// SplitSlice splits slice to several batches with size equals batchSize (except last)
func SplitSlice(input []int, batchSize int) ([][]int, error) {
	if input == nil {
		return nil, nil
	}

	if batchSize < 1 {
		return nil, errors.New("batchSize must be greater than 0")
	}

	batchedCapacity := (len(input) + batchSize - 1) / batchSize
	batched := make([][]int, 0, batchedCapacity)

	for i := 0; i < len(input); i += batchSize {
		rightBound := i + batchSize

		if rightBound > len(input) {
			rightBound = len(input)
		}
		batched = append(batched, input[i:rightBound])
	}
	return batched, nil
}
