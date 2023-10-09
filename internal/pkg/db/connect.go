package db

import (
	"context"
	"fmt"
	"iot-device-tracker-service/internal/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

func getPostgresDSN() string {
	cfg := config.GetPostgresConfig()
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresDbHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDb)
}

func Connect(ctx context.Context) (*PgxDB, error) {
	postgresDSN := getPostgresDSN()
	pool, err := pgxpool.Connect(ctx, postgresDSN)
	if err != nil {
		return nil, err
	}

	return NewDatabase(pool), nil
}
