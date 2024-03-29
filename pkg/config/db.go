package config

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDatabaseConnectionPool(ctx context.Context) *pgxpool.Pool {
	conn, err := pgxpool.New(ctx, "postgres://admin:admin@localhost:5432/ipdb")
	if err != nil {
		panic(err)
	}

	return conn
}
