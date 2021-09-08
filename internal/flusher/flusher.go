package flusher

import (
	"context"
	"github.com/opentracing/opentracing-go"
	opentracingLog "github.com/opentracing/opentracing-go/log"
	"github.com/ozonva/ova-plan-api/internal/models"
	"github.com/ozonva/ova-plan-api/internal/repo"
	"github.com/ozonva/ova-plan-api/internal/utils"
	"github.com/ozonva/ova-plan-api/internal/utils/tracing"
	"log"
)

// Flusher - interface for saving plans into repo
type Flusher interface {
	// Flush save plans into repo. Return plans which could not be saved
	Flush(ctx context.Context, plans []models.Plan) []models.Plan
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

func (f *flusher) Flush(ctx context.Context, plans []models.Plan) []models.Plan {
	span := tracing.StartChildSpan(ctx, "Flush plans")
	defer span.Finish()
	span.LogFields(opentracingLog.Int("plans count", len(plans)))
	batched, err := utils.SplitSlicePlan(plans, f.chunkSize)

	if err != nil {
		log.Printf("Batching error: %v", err.Error())
		return plans
	}

	failed := make([]models.Plan, 0)
	for _, batch := range batched {
		childSpan := opentracing.StartSpan("Flush batch", opentracing.ChildOf(span.Context()))
		childSpan.LogFields(opentracingLog.Int("plans count", len(plans)))
		err := f.planRepo.AddEntities(ctx, batch)
		if err != nil {
			childSpan.LogFields(opentracingLog.String("error", err.Error()))
			childSpan.Finish()
			log.Printf("Batch saving error, skip batch. %v", err)
			failed = append(failed, batch...)
		}
		childSpan.Finish()
	}

	if len(failed) == 0 {
		return nil
	}

	log.Printf("Some entities has not saved: %v", failed)
	return failed
}
