package db

import (
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-plan-api/internal/config"
	"github.com/rs/zerolog/log"
)

func Connect(conf config.DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", conf.GetDsn())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	log.Info().Msg("Db connected")
	return db, nil
}
