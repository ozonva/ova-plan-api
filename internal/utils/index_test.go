package utils

import (
	"errors"
	"github.com/ozonva/ova-plan-api/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIndexPlan(t *testing.T) {
	plans := []models.Plan{
		*models.NewPlan(0, 123, "Сдать задание 3", "Оно ждёт", time.Now(), time.Date(2021, 8, 16, 0, 0, 0, 0, time.UTC)),
		*models.NewPlan(1, 123, "Пройти курс до конца", "", time.Now(), time.Date(2021, 9, 17, 0, 0, 0, 0, time.UTC)),
		*models.NewPlan(42, 123, "Съездить на море", "Красивое", time.Now(), time.Date(2022, 7, 9, 0, 0, 0, 0, time.UTC)),
		*models.NewPlan(42, 123, "Научиться рисовать акварелью", "Сложно", time.Now(), time.Date(2023, 7, 9, 0, 0, 0, 0, time.UTC)),
	}

	tables := []struct {
		input          []models.Plan
		expectedResult map[uint64]models.Plan
		expectedError  error
	}{
		{[]models.Plan{plans[0]}, map[uint64]models.Plan{0: plans[0]}, nil},
		{[]models.Plan{plans[0], plans[1], plans[2]}, map[uint64]models.Plan{0: plans[0], 1: plans[1], 42: plans[2]}, nil},
		{[]models.Plan{}, map[uint64]models.Plan{}, nil},
		{nil, nil, InputSliceIsNil},
		{[]models.Plan{plans[2], plans[3]}, nil, IdentifiersNotUnique},
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
