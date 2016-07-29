package models

import (
	"time"

	"github.com/gernest/wuxia/db"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUser(store *db.DB, u *User) error {
	u.Name = Sanitize(u.Name)
	return nil
}
