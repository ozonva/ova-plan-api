package service

import (
	"context"
	"github.com/ozonva/ova-plan-api/internal/models"
	"github.com/ozonva/ova-plan-api/internal/repo"
	api "github.com/ozonva/ova-plan-api/pkg/ova-plan-api/github.com/ozonva/ova-plan-api/pkg/ova-plan-api"
	"github.com/rs/zerolog/log"
	"time"
)

// implements PlanApiServer
type planApiService struct {
	api.UnimplementedPlanApiServer
	planRepo *repo.PlanRepo
}

func (s *planApiService) CreatePlan(ctx context.Context, request *api.CreatePlanRequest) (*api.CreatePlanResponse, error) {
	log.Info().
		Str("call grpc method", "CreatePlan").
		Str("request", request.String()).
		Send()

	id, err := (*s.planRepo).AddEntity(models.NewPlan(
		0,
		request.Plan.UserId,
		request.Plan.Title,
		request.Plan.Description,
		time.Now(),
		request.Plan.DeadlineAt.AsTime()))
	if err != nil {
		return nil, err
	}

	return &api.CreatePlanResponse{PlanId: id}, nil
}

func (s *planApiService) DescribePlan(ctx context.Context, request *api.DescribePlanRequest) (*api.DescribePlanResponse, error) {
	log.Info().
		Str("call grpc method", "DescribePlan").
		Str("request", request.String()).
		Send()
	return &api.DescribePlanResponse{}, nil
}

func (s *planApiService) ListPlans(ctx context.Context, request *api.ListPlansRequest) (*api.ListPlansResponse, error) {
	log.Info().
		Str("call grpc method", "ListPlans").
		Str("request", request.String()).
		Send()
	return &api.ListPlansResponse{}, nil
}

func (s *planApiService) RemovePlan(ctx context.Context, request *api.RemovePlanRequest) (*api.RemovePlanResponse, error) {
	log.Info().
		Str("call grpc method", "RemovePlan").
		Str("request", request.String()).
		Send()
	return &api.RemovePlanResponse{}, nil
}

func New(planRepo *repo.PlanRepo) api.PlanApiServer {
	return &planApiService{planRepo: planRepo}
}
