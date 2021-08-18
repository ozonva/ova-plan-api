package utils

import (
	"errors"
	"github.com/ozonva/ova-plan-api/internal/plan"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIndexPlan(t *testing.T) {
	plans := []plan.Plan{
		*plan.NewPlan(0, 123, "Сдать задание 3", "Оно ждёт", time.Now(), time.Date(2021, 8, 16, 0, 0, 0, 0, time.UTC)),
		*plan.NewPlan(1, 123, "Пройти курс до конца", "", time.Now(), time.Date(2021, 9, 17, 0, 0, 0, 0, time.UTC)),
		*plan.NewPlan(42, 123, "Съездить на море", "Красивое", time.Now(), time.Date(2022, 7, 9, 0, 0, 0, 0, time.UTC)),
		*plan.NewPlan(42, 123, "Научиться рисовать акварелью", "Сложно", time.Now(), time.Date(2023, 7, 9, 0, 0, 0, 0, time.UTC)),
	}

	tables := []struct {
		input          []plan.Plan
		expectedResult map[uint64]plan.Plan
		expectedError  error
	}{
		{[]plan.Plan{plans[0]}, map[uint64]plan.Plan{0: plans[0]}, nil},
		{[]plan.Plan{plans[0], plans[1], plans[2]}, map[uint64]plan.Plan{0: plans[0], 1: plans[1], 42: plans[2]}, nil},
		{[]plan.Plan{}, map[uint64]plan.Plan{}, nil},
		{nil, nil, InputSliceIsNil},
		{[]plan.Plan{plans[2], plans[3]}, nil, IdentifiersNotUnique},
	}

	for _, table := range tables {
		indexPlan, err := IndexPlan(table.input)
		if table.expectedError == nil {
			assert.Equal(t, indexPlan, table.expectedResult)
		} else {
			assert.True(t, errors.Is(err, table.expectedError))
		}
	}

}
