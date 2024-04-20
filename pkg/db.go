package pkg

import (
	"context"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RunDatabaseSchemaMigration() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://admin:admin@localhost:5432/ipdb?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}

func InitDatabaseConnectionPool(ctx context.Context) *pgxpool.Pool {
	conn, err := pgxpool.New(ctx, "postgres://admin:admin@localhost:5432/ipdb")
	if err != nil {
		panic(err)
	}

	return conn
}
