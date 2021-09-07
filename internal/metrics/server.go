package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

const metricsServerAddress = ":8090"

func RunServer() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(metricsServerAddress, nil); err != nil {
			log.Fatal().Err(err).Msgf("An error occured on starting prometheus server, %s", err)
		}
		log.Info().Msgf("metrics server started on %", metricsServerAddress)
	}()
}
