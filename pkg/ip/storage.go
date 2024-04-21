package ip

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

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

func (s *Storage) CreateRequestIPEntity(entity RequestIPEntity) (RequestIPEntity, error) {
	_, err := s.conn.Exec(
		s.ctx,
		`INSERT INTO request_ips (id, ip, created_at) VALUES ($1, $2, $3);`,
		entity.ID,
		entity.IP,
		entity.CreatedAt,
	)
	if err != nil {
		return RequestIPEntity{}, err
	}

	return entity, nil
}

func (s *Storage) GetStatsIPEntities() ([]StatsIPEntity, error) {
	rows, err := s.conn.Query(
		s.ctx,
		"SELECT ip, COUNT(*) as count FROM request_ips GROUP BY ip ORDER BY count DESC LIMIT 10;",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []StatsIPEntity
	for rows.Next() {
		var entity StatsIPEntity
		if err := rows.Scan(&entity.IP, &entity.Count); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return entities, nil
}
