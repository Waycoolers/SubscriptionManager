package db

import (
	"context"
	"fmt"
	"log/slog"
	"subscriptionmanager/internal/infrastructure/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, cfg *config.DBConfig) (*DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}
	slog.Info("Connected to database successfully")

	return &DB{Pool: pool}, nil
}

func Close(pool *pgxpool.Pool) {
	slog.Info("Closing database connection...")
	if pool != nil {
		pool.Close()
	}
}
