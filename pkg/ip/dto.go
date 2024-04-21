package ip

import "time"

type RequestIPResponseDTO struct {
	ID        string    `json:"id"`
	IP        string    `json:"ip"`
	CreatedAt time.Time `json:"created_at"`
}

type StatsIPResponseDTO struct {
	IP    string `json:"ip"`
	Count int    `json:"count"`
}
