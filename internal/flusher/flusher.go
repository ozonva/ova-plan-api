package flusher

import (
	"github.com/ozonva/ova-plan-api/internal/models"
	"github.com/ozonva/ova-plan-api/internal/repo"
	"github.com/ozonva/ova-plan-api/internal/utils"
)

// Flusher - interface for saving plans into repo
type Flusher interface {
	// Flush save plans into repo. Return plans which could not be saved
	Flush(plans []models.Plan) []models.Plan
}

// NewFlusher creates Flusher with batching save support
func NewFlusher(
	chunkSize int,
	planRepo repo.PlanRepo,
) Flusher {
	return &flusher{
		chunkSize: chunkSize,
		planRepo:  planRepo,
	}
}

type flusher struct {
	chunkSize int
	planRepo  repo.PlanRepo
}

func (f *flusher) Flush(plans []models.Plan) []models.Plan {
	batched, err := utils.SplitSlicePlan(plans, f.chunkSize)

	if err != nil {
		return plans
	}

	failed := make([]models.Plan, 0)

	for _, batch := range batched {
		err := f.planRepo.AddEntities(batch)
		if err != nil {
			failed = append(failed, batch...)
		}
	}

	if len(failed) == 0 {
		return nil
	}
	return failed
}
