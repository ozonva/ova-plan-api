package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-plan-api/internal/models"
)

// PlanRepo - interface for storage Plan entities
type PlanRepo interface {
	AddEntity(ctx context.Context, entity *models.Plan) (uint64, error)
	AddEntities(ctx context.Context, entities []models.Plan) error
	ListEntities(ctx context.Context, limit, offset uint64) ([]models.Plan, error)
	DescribeEntity(ctx context.Context, entityId uint64) (*models.Plan, error)
	RemoveEntity(ctx context.Context, entityId uint64) error
}

type planRepo struct {
	db *sqlx.DB
}

func (p *planRepo) AddEntity(ctx context.Context, entity *models.Plan) (uint64, error) {
	query := squirrel.Insert("plans").
		Columns("user_id", "title", "description", "created_at", "deadline_at").
		Values(entity.UserId, entity.Title, entity.Description, entity.CreatedAt, entity.DeadlineAt).
		Suffix("RETURNING id").
		RunWith(p.db).
		PlaceholderFormat(squirrel.Dollar)

	var id uint64
	err := query.ScanContext(ctx, &id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *planRepo) AddEntities(ctx context.Context, entities []models.Plan) error {
	query := squirrel.Insert("plans").
		Columns("user_id", "title", "description", "created_at", "deadline_at").
		RunWith(p.db).
		PlaceholderFormat(squirrel.Dollar)

	for _, entity := range entities {
		query = query.Values(entity.UserId, entity.Title, entity.Description, entity.CreatedAt, entity.DeadlineAt)
	}

	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *planRepo) ListEntities(ctx context.Context, limit, offset uint64) ([]models.Plan, error) {
	query := squirrel.Select("id", "user_id", "title", "description", "created_at", "deadline_at").
		From("plans").
		Limit(limit).
		Offset(offset).
		OrderBy("id").
		RunWith(p.db).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := query.QueryContext(ctx)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]models.Plan, 0, 0)

	for rows.Next() {
		plan, err := readRow(rows)

		if err != nil {
			return nil, err
		}
		result = append(result, *plan)
	}
	return result, nil
}

func readRow(rows *sql.Rows) (*models.Plan, error) {
	plan := models.NewEmptyPlan()
	err := rows.Scan(
		&plan.Id,
		&plan.UserId,
		&plan.Title,
		&plan.Description,
		&plan.CreatedAt,
		&plan.DeadlineAt)

	if err != nil {
		return nil, err
	}
	return plan, nil
}

func (p *planRepo) DescribeEntity(ctx context.Context, entityId uint64) (*models.Plan, error) {
	query := squirrel.Select("id", "user_id", "title", "description", "created_at", "deadline_at").
		From("plans").
		Where(squirrel.Eq{"id": entityId}).
		RunWith(p.db).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("there is no plan with id %v", entityId)
	}

	plan, err := readRow(rows)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (p *planRepo) RemoveEntity(ctx context.Context, entityId uint64) error {
	_, err := squirrel.Delete("plans").
		RunWith(p.db).
		Where(squirrel.Eq{"id": entityId}).
		PlaceholderFormat(squirrel.Dollar).
		ExecContext(ctx)
	return err
}

func New(db *sqlx.DB) PlanRepo {
	return &planRepo{db: db}
}
