package repo

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-plan-api/internal/models"
)

// PlanRepo - interface for storage Plan entities
type PlanRepo interface {
	AddEntity(entity *models.Plan) (uint64, error)
	AddEntities(entities []models.Plan) error
	ListEntities(limit, offset uint64) ([]models.Plan, error)
	DescribeEntity(entityId uint64) (*models.Plan, error)
}

type planRepo struct {
	db *sqlx.DB
}

func (p *planRepo) AddEntity(entity *models.Plan) (uint64, error) {
	query := squirrel.Insert("plans").
		Columns("user_id", "title", "description", "created_at", "deadline_at").
		Values(entity.UserId, entity.Title, entity.Description, entity.CreatedAt, entity.DeadlineAt).
		Suffix("RETURNING id").
		RunWith(p.db).
		PlaceholderFormat(squirrel.Dollar)

	var id uint64
	err := query.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *planRepo) AddEntities(entities []models.Plan) error {
	panic("implement me")
}

func (p *planRepo) ListEntities(limit, offset uint64) ([]models.Plan, error) {
	panic("implement me")
}

func (p *planRepo) DescribeEntity(entityId uint64) (*models.Plan, error) {
	panic("implement me")
}

func New(db *sqlx.DB) PlanRepo {
	return &planRepo{db: db}
}
