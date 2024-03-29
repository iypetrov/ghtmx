package ip

import "time"

type RequestIPEntity struct {
	ID        string    `db:"id"`
	IP        string    `db:"ip"`
	CreatedOn time.Time `db:"created_on"`
}
