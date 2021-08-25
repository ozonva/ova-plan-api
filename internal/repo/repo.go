package repo

import "github.com/ozonva/ova-plan-api/internal/models"

// PlanRepo - interface for storage Plan entities
type PlanRepo interface {
	AddEntities(entities []models.Plan) error
	ListEntities(limit, offset uint64) ([]models.Plan, error)
	DescribeEntity(entityId uint64) (*models.Plan, error)
}
