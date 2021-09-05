package server

import (
	api "github.com/ozonva/ova-plan-api/pkg/ova-plan-api/github.com/ozonva/ova-plan-api/pkg/ova-plan-api"
	zerolog "github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server interface {
	Run(port string) error
	Stop() error
}

// implements Server
type grpcServer struct {
	planApiService *api.PlanApiServer
}

func (g *grpcServer) Stop() error {
	panic("implement me")
}

func (g *grpcServer) Run(port string) error {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	api.RegisterPlanApiServer(s, *g.planApiService)
	zerolog.Info().Msg("Grpc server started")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func New(planApiService *api.PlanApiServer) Server {
	return &grpcServer{planApiService: planApiService}
}
