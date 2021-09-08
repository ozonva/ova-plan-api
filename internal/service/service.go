package service

import (
	"context"
	"github.com/opentracing/opentracing-go"
	planFlusher "github.com/ozonva/ova-plan-api/internal/flusher"
	"github.com/ozonva/ova-plan-api/internal/kafka"
	"github.com/ozonva/ova-plan-api/internal/metrics"
	"github.com/ozonva/ova-plan-api/internal/models"
	"github.com/ozonva/ova-plan-api/internal/repo"
	"github.com/ozonva/ova-plan-api/internal/utils/tracing"
	api "github.com/ozonva/ova-plan-api/pkg/ova-plan-api/github.com/ozonva/ova-plan-api/pkg/ova-plan-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// implements PlanApiServer
type planApiService struct {
	api.UnimplementedPlanApiServer
	planRepo      repo.PlanRepo
	flusher       planFlusher.Flusher
	kafkaProducer kafka.Producer
}

func (s *planApiService) CreatePlan(ctx context.Context, request *api.CreatePlanRequest) (*api.CreatePlanResponse, error) {
	span := opentracing.StartSpan("CreatePlan rpc")
	defer span.Finish()

	log.Info().
		Str("call grpc method", "CreatePlan").
		Str("request", request.String()).
		Send()

	plan := newPlan(request.Plan)
	id, err := s.planRepo.AddEntity(tracing.CtxWithParentSpan(ctx, span), plan)
	if err != nil {
		return nil, err
	}
	msgs, err := kafka.NewCreatePlanMessages([]models.Plan{*plan})
	if err != nil {
		return nil, err
	}
	err = s.kafkaProducer.Send(msgs)
	if err != nil {
		return nil, err
	}

	metrics.AddCreatePlanSucceeds(1)

	return &api.CreatePlanResponse{PlanId: id}, nil
}

func (s *planApiService) MultiCreatePlan(ctx context.Context, request *api.MultiCreatePlanRequest) (*api.MultiCreatePlanResponse, error) {
	span := opentracing.StartSpan("MultiCreatePlan rpc")
	defer span.Finish()

	log.Info().
		Str("call grpc method", "MultiCreatePlan").
		Str("request", request.String()).
		Send()

	planModels := make([]models.Plan, 0, len(request.GetPlans()))

	for _, plan := range request.GetPlans() {
		planModels = append(planModels, *newPlan(plan))
	}

	notCreated := s.flusher.Flush(tracing.CtxWithParentSpan(ctx, span), planModels)

	msgs, err := kafka.NewCreatePlanMessages(planModels)
	if err != nil {
		return nil, err
	}
	err = s.kafkaProducer.Send(msgs)
	if err != nil {
		return nil, err
	}

	metrics.AddCreatePlanSucceeds(len(request.GetPlans()) - len(notCreated))

	return &api.MultiCreatePlanResponse{}, nil
}

func (s *planApiService) DescribePlan(ctx context.Context, request *api.DescribePlanRequest) (*api.DescribePlanResponse, error) {
	span := opentracing.StartSpan("DescribePlan rpc")
	defer span.Finish()

	log.Info().
		Str("call grpc method", "DescribePlan").
		Str("request", request.String()).
		Send()

	plan, err := s.planRepo.DescribeEntity(tracing.CtxWithParentSpan(ctx, span), request.PlanId)
	if err != nil {
		return nil, err
	}

	protoPlan := mapPlanToProto(plan)

	return &api.DescribePlanResponse{Plan: protoPlan}, nil
}

func (s *planApiService) ListPlans(ctx context.Context, request *api.ListPlansRequest) (*api.ListPlansResponse, error) {
	span := opentracing.StartSpan("ListPlans rpc")
	defer span.Finish()

	log.Info().
		Str("call grpc method", "ListPlans").
		Str("request", request.String()).
		Send()

	plans, err := s.planRepo.ListEntities(tracing.CtxWithParentSpan(ctx, span), request.GetLimit()+1, request.GetOffset())
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
	span := opentracing.StartSpan("RemovePlan rpc")
	defer span.Finish()

	log.Info().
		Str("call grpc method", "RemovePlan").
		Str("request", request.String()).
		Send()

	err := s.planRepo.RemoveEntity(tracing.CtxWithParentSpan(ctx, span), request.PlanId)
	if err != nil {
		return &api.RemovePlanResponse{Error: err.Error()}, nil
	}

	msgs, err := kafka.NewRemovePlanMessages([]uint64{request.PlanId})
	if err != nil {
		return nil, err
	}
	err = s.kafkaProducer.Send(msgs)
	if err != nil {
		return nil, err
	}

	metrics.AddRemovePlanSucceeds(1)

	return &api.RemovePlanResponse{}, nil
}

func (s *planApiService) UpdatePlan(ctx context.Context, request *api.UpdatePlanRequest) (*api.UpdatePlanResponse, error) {
	span := opentracing.StartSpan("UpdatePlan rpc")
	defer span.Finish()

	log.Info().
		Str("call grpc method", "UpdatePlan").
		Str("request", request.String()).
		Send()

	plan := newPlan(request.GetPlan())
	err := s.planRepo.UpdateEntity(ctx, request.PlanId, plan)
	if err != nil {
		return nil, err
	}

	msgs, err := kafka.NewUpdatePlanMessages([]models.Plan{*plan})
	if err != nil {
		return nil, err
	}
	err = s.kafkaProducer.Send(msgs)
	if err != nil {
		return nil, err
	}

	metrics.AddUpdatePlanSucceeds(1)

	return &api.UpdatePlanResponse{}, nil
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

func newPlan(planTemplate *api.PlanTemplate) *models.Plan {
	return models.NewPlan(
		0,
		planTemplate.UserId,
		planTemplate.Title,
		planTemplate.Description,
		time.Now(),
		planTemplate.DeadlineAt.AsTime())
}

func New(planRepo *repo.PlanRepo, flusher *planFlusher.Flusher, kafkaProducer *kafka.Producer) api.PlanApiServer {
	return &planApiService{planRepo: *planRepo, flusher: *flusher, kafkaProducer: *kafkaProducer}
}
