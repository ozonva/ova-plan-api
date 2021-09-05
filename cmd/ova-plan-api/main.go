package main

import (
	"github.com/ozonva/ova-plan-api/internal/config"
	database "github.com/ozonva/ova-plan-api/internal/db"
	"github.com/ozonva/ova-plan-api/internal/repo"
	"github.com/ozonva/ova-plan-api/internal/server"
	"github.com/ozonva/ova-plan-api/internal/service"
	"log"
)

func main() {
	dbConfig := config.NewEnvVarDatabaseConfig()
	db, err := database.Connect(dbConfig)
	if err != nil {
		log.Fatalf("Database connect failed, %v", err.Error())
	}
	defer db.Close()

	planRepo := repo.New(db)
	planApiService := service.New(&planRepo)
	grpcServer := server.New(&planApiService)

	err = grpcServer.Run(":8080")
	if err != nil {
		log.Fatalf("Grpc start failed, %v", err.Error())
	}
}
