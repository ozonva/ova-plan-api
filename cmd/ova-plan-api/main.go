package main

import (
	"github.com/opentracing/opentracing-go"
	"github.com/ozonva/ova-plan-api/internal/config"
	database "github.com/ozonva/ova-plan-api/internal/db"
	"github.com/ozonva/ova-plan-api/internal/flusher"
	"github.com/ozonva/ova-plan-api/internal/kafka"
	"github.com/ozonva/ova-plan-api/internal/metrics"
	"github.com/ozonva/ova-plan-api/internal/repo"
	"github.com/ozonva/ova-plan-api/internal/server"
	"github.com/ozonva/ova-plan-api/internal/service"
	"github.com/ozonva/ova-plan-api/internal/tracing"
	"github.com/rs/zerolog/log"
)

func main() {
	// Tracer
	tracer, traceCloser := tracing.InitTracer()
	opentracing.SetGlobalTracer(tracer)
	defer traceCloser.Close()

	// Kafka
	kafkaConfig := config.NewKafkaConfig()
	kafkaProducer := kafka.NewSyncProducer(kafkaConfig)

	// DB
	dbConfig := config.NewEnvVarDatabaseConfig()
	db, err := database.Connect(dbConfig)
	if err != nil {
		log.Fatal().Err(err).Msgf("Database connect failed")
	}
	defer db.Close()

	// Metrics
	metrics.RunServer()

	// App
	planRepo := repo.New(db)
	planFlusher := flusher.NewFlusher(2, planRepo)
	planApiService := service.New(&planRepo, &planFlusher, &kafkaProducer)
	grpcServer := server.New(&planApiService)

	err = grpcServer.Run(":8080")
	if err != nil {
		log.Fatal().Err(err).Msgf("Grpc start failed")
	}
}
