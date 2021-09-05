package service

import (
	"context"
	"github.com/ozonva/ova-plan-api/internal/models"
	"github.com/ozonva/ova-plan-api/internal/repo"
	planSaver "github.com/ozonva/ova-plan-api/internal/saver"
	api "github.com/ozonva/ova-plan-api/pkg/ova-plan-api/github.com/ozonva/ova-plan-api/pkg/ova-plan-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// implements PlanApiServer
type planApiService struct {
	api.UnimplementedPlanApiServer
	planRepo repo.PlanRepo
	saver    planSaver.Saver
}

func (s *planApiService) CreatePlan(ctx context.Context, request *api.CreatePlanRequest) (*api.CreatePlanResponse, error) {
	log.Info().
		Str("call grpc method", "CreatePlan").
		Str("request", request.String()).
		Send()

	id, err := s.planRepo.AddEntity(newPlan(request.Plan))
	if err != nil {
		return nil, err
	}

	return &api.CreatePlanResponse{PlanId: id}, nil
}

func (s *planApiService) MultiCreatePlan(ctx context.Context, request *api.MultiCreatePlanRequest) (*api.MultiCreatePlanResponse, error) {
	log.Info().
		Str("call grpc method", "MultiCreatePlan").
		Str("request", request.String()).
		Send()

	for _, plan := range request.GetPlans() {
		planModel := newPlan(plan)
		s.saver.Save(*planModel)
	}

	return &api.MultiCreatePlanResponse{}, nil
}

func newPlan(planTemplate *api.PlanTemplate) *models.Plan {
	return models.NewPlan(
		0,
		planTemplate.UserId,
		planTemplate.Title,
		planTemplate.Description,
		time.Now(),
		planTemplate.DeadlineAt.AsTime())
}

func (s *planApiService) DescribePlan(ctx context.Context, request *api.DescribePlanRequest) (*api.DescribePlanResponse, error) {
	log.Info().
		Str("call grpc method", "DescribePlan").
		Str("request", request.String()).
		Send()

	plan, err := s.planRepo.DescribeEntity(request.PlanId)
	if err != nil {
		return nil, err
	}

	protoPlan := mapPlanToProto(plan)

	return &api.DescribePlanResponse{Plan: protoPlan}, nil
}

func (s *planApiService) ListPlans(ctx context.Context, request *api.ListPlansRequest) (*api.ListPlansResponse, error) {
	log.Info().
		Str("call grpc method", "ListPlans").
		Str("request", request.String()).
		Send()

	plans, err := s.planRepo.ListEntities(request.GetLimit()+1, request.GetOffset())
	if err != nil {
		return nil, err
	}

	hasMore := false
	resultLength := len(plans)
	if resultLength > int(request.GetLimit()) {
		resultLength = int(request.GetLimit())
		hasMore = true
	}

	protoPlans := make([]*api.Plan, 0, len(plans))
	for _, plan := range plans {
		if len(protoPlans) == int(request.GetLimit()) {
			break
		}
		protoPlans = append(protoPlans, mapPlanToProto(&plan))
	}

	return &api.ListPlansResponse{
		Plans:   protoPlans,
		HasMore: hasMore,
	}, nil
}

func (s *planApiService) RemovePlan(ctx context.Context, request *api.RemovePlanRequest) (*api.RemovePlanResponse, error) {
	log.Info().
		Str("call grpc method", "RemovePlan").
		Str("request", request.String()).
		Send()

	err := s.planRepo.RemoveEntity(request.PlanId)
	if err != nil {
		return &api.RemovePlanResponse{Error: err.Error()}, nil
	}

	return &api.RemovePlanResponse{}, nil
}

func mapPlanToProto(plan *models.Plan) *api.Plan {
	return &api.Plan{
		PlanId:      plan.Id,
		UserId:      plan.UserId,
		Title:       plan.Title,
		Description: plan.Description,
		CreatedAt:   timestamppb.New(plan.CreatedAt),
		DeadlineAt:  timestamppb.New(plan.DeadlineAt),
	}
}

func New(planRepo *repo.PlanRepo, saver *planSaver.Saver) api.PlanApiServer {
	return &planApiService{planRepo: *planRepo, saver: *saver}
}
