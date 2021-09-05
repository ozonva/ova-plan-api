package main

import (
	"github.com/opentracing/opentracing-go"
	"github.com/ozonva/ova-plan-api/internal/config"
	database "github.com/ozonva/ova-plan-api/internal/db"
	"github.com/ozonva/ova-plan-api/internal/flusher"
	"github.com/ozonva/ova-plan-api/internal/repo"
	"github.com/ozonva/ova-plan-api/internal/server"
	"github.com/ozonva/ova-plan-api/internal/service"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"io"
)

func main() {
	tracer, traceCloser := initTracer()
	opentracing.SetGlobalTracer(tracer)
	defer traceCloser.Close()

	dbConfig := config.NewEnvVarDatabaseConfig()
	db, err := database.Connect(dbConfig)

	if err != nil {
		log.Fatal().Msgf("Database connect failed, %v", err.Error())
	}
	defer db.Close()

	planRepo := repo.New(db)
	planFlusher := flusher.NewFlusher(2, planRepo)
	if err != nil {
		log.Fatal().Msgf("Can't create new plan flusher, %v", err.Error())
	}
	planApiService := service.New(&planRepo, &planFlusher)
	grpcServer := server.New(&planApiService)

	err = grpcServer.Run(":8080")
	if err != nil {
		log.Fatal().Msgf("Grpc start failed, %v", err.Error())
	}
}

func initTracer() (opentracing.Tracer, io.Closer) {
	cfg := jaegercfg.Configuration{
		ServiceName: "ova-plan-api",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		log.Fatal().Msgf("Can't create tracer, %v", err)
	}
	return tracer, closer
}
