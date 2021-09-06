package config

import (
	"fmt"
	"os"
)

type envVarDatabaseConfig struct {
}

func (e *envVarDatabaseConfig) GetDsn() string {
	user := os.Getenv("OVA_PLAN_DB_USER")
	password := os.Getenv("OVA_PLAN_DB_PASSWORD")
	host := os.Getenv("OVA_PLAN_DB_HOST")
	port := os.Getenv("OVA_PLAN_DB_PORT")
	dbName := os.Getenv("OVA_PLAN_DB_NAME")

	return fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbName)
}

func NewEnvVarDatabaseConfig() DatabaseConfig {
	return &envVarDatabaseConfig{}
}
