package main

import (
	database "github.com/ozonva/ova-plan-api/internal/db"
	"github.com/ozonva/ova-plan-api/internal/repo"
	"github.com/ozonva/ova-plan-api/internal/server"
	"github.com/ozonva/ova-plan-api/internal/service"
)

func main() {
	//TODO: To config
	db, err := database.Connect("postgresql://ova_plan:ova_plan@127.0.0.1:5432/ova_plan_db?sslmode=disable")
	if err != nil {
		panic("Database connect failed") // TODO
	}
	defer db.Close()

	planRepo := repo.New(db)
	planApiService := service.New(&planRepo)
	grpcServer := server.New(&planApiService)

	err = grpcServer.Run(":8080")
	if err != nil {
		panic("Grpc start failed") // TODO
	}
}
