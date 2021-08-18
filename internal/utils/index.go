package utils

import (
	"errors"
	"fmt"
	"github.com/ozonva/ova-plan-api/internal/plan"
)

var (
	InputSliceIsNil      = errors.New("input slice must be initialized")
	IdentifiersNotUnique = errors.New("identifiers not unique")
)

func IndexPlan(plans []plan.Plan) (map[uint64]plan.Plan, error) {
	if plans == nil {
		return nil, InputSliceIsNil
	}

	index := make(map[uint64]plan.Plan, len(plans))

	for _, pl := range plans {
		if _, exist := index[pl.Id]; exist {
			return nil, fmt.Errorf("%w: plan with identifier %v occurs more than once", IdentifiersNotUnique, pl.Id)
		}
		index[pl.Id] = pl
	}
	return index, nil
}
