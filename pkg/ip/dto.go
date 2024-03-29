package ip

import "time"

type RequestIPResponseDTO struct {
	ID        string    `json:"id"`
	IP        string    `json:"ip"`
	CreatedOn time.Time `json:"created_on"`
}
