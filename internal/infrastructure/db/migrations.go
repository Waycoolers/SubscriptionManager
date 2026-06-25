package db

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log/slog"
	"subscriptionmanager/internal/infrastructure/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func Migrate(cfg *config.DBConfig) error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			slog.Error("Failed to close database connection")
		}
	}(db)

	slog.Info("Migrating database")
	files, err := iofs.New(migrationFiles, ".")
	if err != nil {
		return err
	}
	driver, err := pgx.WithInstance(db, &pgx.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("iofs", files, "pgx", driver)
	if err != nil {
		return err
	}

	err = m.Up()

	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
		slog.Info("Migrations haven't changed anything")
	}

	slog.Info("Migrations complete")
	return nil
}
