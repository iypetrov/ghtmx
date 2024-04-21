package ip

import (
	"time"

	"github.com/google/uuid"
)

type RequestIPModel struct {
	ID        uuid.UUID
	IP        string
	CreatedAt time.Time
}

type StatsIPModel struct {
	IP    string
	Count int
}
