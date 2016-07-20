package models

import "time"

type Session struct {
	ID        int64
	Data      []byte
	CreatedOn time.Time
	UpdatedOn timeTime
	ExpiresOn time.Time
}
