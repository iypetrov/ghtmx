package ip

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetStatus() error
}

type Storage struct {
	ctx  context.Context
	conn *pgxpool.Pool
}

func NewStorage(ctx context.Context, conn *pgxpool.Pool) *Storage {
	return &Storage{
		ctx:  ctx,
		conn: conn,
	}
}

func (s *Storage) GetStatus() (string, error) {
	var status string
	err := s.conn.QueryRow(s.ctx, "select 'Hello, world!'").Scan(&status)
	if err != nil {
		return "", err
	}

	return status, nil
}
