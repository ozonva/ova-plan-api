package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	createPlanSucceeds = promauto.NewCounter(prometheus.CounterOpts{
		Name: "create_plan_succeed",
		Help: "Number of successful plan creations",
	})

	updatePlanSucceeds = promauto.NewCounter(prometheus.CounterOpts{
		Name: "update_plan_succeed",
		Help: "Number of successful plan updates",
	})

	removePlanSucceeds = promauto.NewCounter(prometheus.CounterOpts{
		Name: "remove_plan_succeed",
		Help: "Number of successful plan removing",
	})
)

func AddCreatePlanSucceeds(count int) {
	createPlanSucceeds.Add(float64(count))
}

func AddUpdatePlanSucceeds(count int) {
	updatePlanSucceeds.Add(float64(count))
}

func AddRemovePlanSucceeds(count int) {
	removePlanSucceeds.Add(float64(count))

}
