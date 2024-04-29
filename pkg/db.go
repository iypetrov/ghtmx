package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/IliyaYavorovPetrov/ghtmx/config"
)

func RunDatabaseSchemaMigration(cfg config.Config) error {
	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("postgres://%s:%s@%s/%s",
			cfg.Storage.Username,
			cfg.Storage.Password,
			cfg.Storage.Addr,
			cfg.Storage.Name,
		))
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	return nil
}

func InitDatabaseConnectionPool(ctx context.Context, cfg config.Config) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(
		ctx,
		fmt.Sprintf(
			"postgres://%s:%s@%s/%s",
			cfg.Storage.Username,
			cfg.Storage.Password,
			cfg.Storage.Addr,
			cfg.Storage.Name,
		),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
