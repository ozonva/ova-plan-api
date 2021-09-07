// +build integration

package integration_test

import (
	"context"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	cl "github.com/ozonva/ova-plan-api/pkg/ova-plan-api/github.com/ozonva/ova-plan-api/pkg/client"
	api "github.com/ozonva/ova-plan-api/pkg/ova-plan-api/github.com/ozonva/ova-plan-api/pkg/ova-plan-api"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
	"time"
)

var _ = Describe("Ova-plan-api integration test", func() {
	var (
		client, _ = cl.NewClient("127.0.0.1", "8080")
	)

	Describe("Positive scenarios", func() {
		It("should create single plan", func() {
			createPlanResponse, err := client.CreatePlan(context.TODO(), &api.CreatePlanRequest{
				Plan: createPlanTemplate(),
			})
			Expect(err).To(BeNil())
			Expect(createPlanResponse.GetPlanId()).To(BeNumerically(">", 0))
		})

		It("should describe created plan", func() {
			planTemplate := createPlanTemplate()
			createPlanResponse, err := client.CreatePlan(context.TODO(), &api.CreatePlanRequest{
				Plan: planTemplate,
			})
			createdPlanId := createPlanResponse.GetPlanId()
			Expect(err).To(BeNil())
			Expect(createdPlanId).To(BeNumerically(">", 0))

			describedPlanResponse, err := client.DescribePlan(context.TODO(), &api.DescribePlanRequest{
				PlanId: createdPlanId,
			})
			describedPlan := describedPlanResponse.GetPlan()
			Expect(err).To(BeNil())
			assertPlan(planTemplate, describedPlan, createdPlanId)
		})

		It("should remove single plan", func() {
			planTemplate := createPlanTemplate()
			createPlanResponse, err := client.CreatePlan(context.TODO(), &api.CreatePlanRequest{
				Plan: planTemplate,
			})
			createdPlanId := createPlanResponse.GetPlanId()
			Expect(err).To(BeNil())
			Expect(createdPlanId).To(BeNumerically(">", 0))

			removePlanResponse, err := client.RemovePlan(context.TODO(), &api.RemovePlanRequest{
				PlanId: createdPlanId,
			})
			Expect(err).To(BeNil())
			Expect(removePlanResponse.GetError()).To(Equal(""))

			describedPlanResponse, _ := client.DescribePlan(context.TODO(), &api.DescribePlanRequest{
				PlanId: createdPlanId,
			})
			Expect(describedPlanResponse.GetPlan()).To(BeNil())
		})

		It("should update single plan", func() {
			planTemplate := createPlanTemplate()
			createPlanResponse, err := client.CreatePlan(context.TODO(), &api.CreatePlanRequest{
				Plan: planTemplate,
			})
			createdPlanId := createPlanResponse.GetPlanId()
			Expect(err).To(BeNil())
			Expect(createdPlanId).To(BeNumerically(">", 0))

			newPlanTemplate := createPlanTemplate()
			_, err = client.UpdatePlan(context.TODO(), &api.UpdatePlanRequest{
				PlanId: createdPlanId,
				Plan:   newPlanTemplate,
			})
			Expect(err).To(BeNil())

			describedPlanResponse, err := client.DescribePlan(context.TODO(), &api.DescribePlanRequest{
				PlanId: createdPlanId,
			})
			Expect(err).To(BeNil())

			assertPlan(newPlanTemplate, describedPlanResponse.GetPlan(), createdPlanId)
		})

		It("should create multiple plans", func() {
			createdFrom := timestamppb.Now()
			plansCountBefore := plansCount(client, createdFrom)

			planTemplates := []*api.PlanTemplate{
				createPlanTemplate(),
				createPlanTemplate(),
				createPlanTemplate(),
			}

			_, err := client.MultiCreatePlan(context.TODO(), &api.MultiCreatePlanRequest{
				Plans: planTemplates,
			})
			Expect(err).To(BeNil())

			plansCountAfter := plansCount(client, createdFrom)
			Expect(plansCountAfter - plansCountBefore).To(Equal(len(planTemplates)))
		})
	})
})

func plansCount(client cl.Client, timestamp *timestamppb.Timestamp) int {
	limit := uint64(10000)
	response, _ := client.ListPlans(context.TODO(), &api.ListPlansRequest{
		CreatedFrom: timestamp,
		Limit:       &limit,
	})

	return len(response.GetPlans())
}

func createPlanTemplate() *api.PlanTemplate {
	return &api.PlanTemplate{
		UserId:      uint64(rand.Intn(100)),
		Title:       fmt.Sprint(rand.Int()),
		Description: fmt.Sprint(rand.Int()),
		DeadlineAt:  timestamppb.Now(),
	}
}

func assertPlan(expected *api.PlanTemplate, actual *api.Plan, expectedId uint64) {
	Expect(actual).To(Not(BeNil()))
	Expect(actual.GetPlanId()).To(Equal(expectedId))
	Expect(actual.GetDescription()).To(Equal(expected.GetDescription()))
	Expect(actual.GetTitle()).To(Equal(expected.GetTitle()))
	Expect(actual.GetDeadlineAt().AsTime().Truncate(time.Second)).
		To(Equal(expected.GetDeadlineAt().AsTime().Truncate(time.Second)))
	Expect(actual.GetCreatedAt()).To(Not(BeNil()))
}
