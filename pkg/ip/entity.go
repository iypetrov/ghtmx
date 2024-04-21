package ip

import (
	"time"

	"github.com/google/uuid"
)

type RequestIPEntity struct {
	ID        uuid.UUID `db:"id"`
	IP        string    `db:"ip"`
	CreatedAt time.Time `db:"created_at"`
}

type StatsIPEntity struct {
	IP    string
	Count int
}
