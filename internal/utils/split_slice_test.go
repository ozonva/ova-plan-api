package utils

import (
	"errors"
	"github.com/ozonva/ova-plan-api/internal/plan"
	"reflect"
	"testing"
	"time"
)

func TestSplitSliceInt(t *testing.T) {
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
		result, err := SplitSliceInt(table.inputSlice, table.inputBatchSize)

		if !reflect.DeepEqual(table.expectedErr, err) {
			t.Errorf("\tWrong error! Actual is \"%v\" but \"%v\" expected", err, table.expectedErr)
		}
		if !reflect.DeepEqual(table.expected, result) {
			t.Errorf("\tWrong result! Actual is %v but %v expected", result, table.expected)
		}
	}
}

func TestSplitSlicePlan(t *testing.T) {
	plans := []plan.Plan{
		*plan.NewPlan(0, 123, "Сдать задание 3", "Оно ждёт", time.Now(), time.Date(2021, 8, 16, 0, 0, 0, 0, time.UTC)),
		*plan.NewPlan(2, 123, "Пройти курс до конца", "", time.Now(), time.Date(2021, 9, 17, 0, 0, 0, 0, time.UTC)),
		*plan.NewPlan(3, 123, "Съездить на море", "Красивое", time.Now(), time.Date(2022, 7, 9, 0, 0, 0, 0, time.UTC)),
		*plan.NewPlan(3, 123, "Научиться рисовать акварелью", "Сложно", time.Now(), time.Date(2023, 7, 9, 0, 0, 0, 0, time.UTC)),
	}

	tables := []struct {
		inputSlice     []plan.Plan
		inputBatchSize int
		expected       [][]plan.Plan
		expectedErr    error
	}{
		{[]plan.Plan{plans[0], plans[1], plans[2], plans[3]}, 2, [][]plan.Plan{{plans[0], plans[1]}, {plans[2], plans[3]}}, nil},
		{[]plan.Plan{plans[0], plans[1], plans[2]}, 1, [][]plan.Plan{{plans[0]}, {plans[1]}, {plans[2]}}, nil},
		{[]plan.Plan{plans[0], plans[1], plans[2]}, 3, [][]plan.Plan{{plans[0], plans[1], plans[2]}}, nil},
		{[]plan.Plan{plans[0], plans[1], plans[2]}, 0, nil, errors.New("batchSize must be greater than 0")},
		{[]plan.Plan{plans[0], plans[1]}, 3, [][]plan.Plan{{plans[0], plans[1]}}, nil},
		{[]plan.Plan{plans[0], plans[1], plans[2]}, 2, [][]plan.Plan{{plans[0], plans[1]}, {plans[2]}}, nil},
		{[]plan.Plan{}, 2, [][]plan.Plan{}, nil},
		{nil, 1, nil, nil},
	}

	for _, table := range tables {
		t.Logf("input: %v %v", table.inputSlice, table.inputBatchSize)
		result, err := SplitSlicePlan(table.inputSlice, table.inputBatchSize)

		if !reflect.DeepEqual(table.expectedErr, err) {
			t.Errorf("\tWrong error! Actual is \"%v\" but \"%v\" expected", err, table.expectedErr)
		}
		if !reflect.DeepEqual(table.expected, result) {
			t.Errorf("\tWrong result! Actual is %v but %v expected", result, table.expected)
		}
	}
}
