package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"

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
	p, err := HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = p
	return nil
}

func HashPassword(pass []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
}
