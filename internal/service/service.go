package service

import (
	"context"
	api "github.com/ozonva/ova-plan-api/pkg/ova-plan-api/github.com/ozonva/ova-plan-api/pkg/ova-plan-api"
)

// implements PlanApiServer
type planApiService struct {
	api.UnimplementedPlanApiServer
}

func (s *planApiService) CreatePlan(ctx context.Context, request *api.CreatePlanRequest) (*api.CreatePlanResponse, error) {
	panic("implement me")
}

func (s *planApiService) DescribePlan(ctx context.Context, request *api.DescribePlanRequest) (*api.DescribePlanResponse, error) {
	panic("implement me")
}

func (s *planApiService) ListPlans(ctx context.Context, request *api.ListPlansRequest) (*api.ListPlansResponse, error) {
	panic("implement me")
}

func (s *planApiService) RemovePlan(ctx context.Context, request *api.RemovePlanRequest) (*api.RemovePlanResponse, error) {
	panic("implement me")
}

func New() api.PlanApiServer {
	return &planApiService{}
}
