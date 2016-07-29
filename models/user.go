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
	var query = `
	BEGIN TRANSACTION;
	  INSERT INTO users (username,password,email,created_at)
		VALUES ($1,$2,$3,now());
	COMMIT;
	`
	tx, err := store.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, u.Name, u.Password, u.Email)
	if err != nil {
		return err
	}
	return tx.Commit()
	return nil
}

func HashPassword(pass []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
}
