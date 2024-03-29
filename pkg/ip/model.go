package ip

import "time"

type RequestIPModel struct {
	ID        string
	IP        string
	CreatedOn time.Time
}
