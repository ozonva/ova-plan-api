package client

import (
	"context"
	"fmt"
	ova_plan_api "github.com/ozonva/ova-plan-api/pkg/ova-plan-api/github.com/ozonva/ova-plan-api/pkg/ova-plan-api"
	"google.golang.org/grpc"
)

type Client interface {
	ova_plan_api.PlanApiClient

	Close() error
}

type client struct {
	connection *grpc.ClientConn
	impl       ova_plan_api.PlanApiClient
}

func (c *client) CreatePlan(ctx context.Context, in *ova_plan_api.CreatePlanRequest, opts ...grpc.CallOption) (*ova_plan_api.CreatePlanResponse, error) {
	return c.impl.CreatePlan(ctx, in, opts...)
}

func (c *client) DescribePlan(ctx context.Context, in *ova_plan_api.DescribePlanRequest, opts ...grpc.CallOption) (*ova_plan_api.DescribePlanResponse, error) {
	return c.impl.DescribePlan(ctx, in, opts...)
}

func (c *client) ListPlans(ctx context.Context, in *ova_plan_api.ListPlansRequest, opts ...grpc.CallOption) (*ova_plan_api.ListPlansResponse, error) {
	return c.impl.ListPlans(ctx, in, opts...)
}

func (c *client) RemovePlan(ctx context.Context, in *ova_plan_api.RemovePlanRequest, opts ...grpc.CallOption) (*ova_plan_api.RemovePlanResponse, error) {
	return c.impl.RemovePlan(ctx, in, opts...)
}

func (c *client) MultiCreatePlan(ctx context.Context, in *ova_plan_api.MultiCreatePlanRequest, opts ...grpc.CallOption) (*ova_plan_api.MultiCreatePlanResponse, error) {
	return c.impl.MultiCreatePlan(ctx, in, opts...)
}

func (c *client) UpdatePlan(ctx context.Context, in *ova_plan_api.UpdatePlanRequest, opts ...grpc.CallOption) (*ova_plan_api.UpdatePlanResponse, error) {
	return c.impl.UpdatePlan(ctx, in, opts...)
}

func (c *client) Close() error {
	if c.connection != nil {
		return c.connection.Close()
	}
	return nil
}

func NewClient(host string, port string) (Client, error) {
	address := fmt.Sprintf("%s:%s", host, port)
	connection, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return &client{
		connection: connection,
		impl:       ova_plan_api.NewPlanApiClient(connection),
	}, nil
}
