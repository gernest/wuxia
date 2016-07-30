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
	  INSERT INTO users (username,password,email,created_at,updated_at)
		VALUES ($1,$2,$3,now(),now());
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

func VerifyPass(hash, pass []byte) error {
	return bcrypt.CompareHashAndPassword(hash, pass)
}

func FindUserByEmail(store *db.DB, email string) (*User, error) {
	var query = `
	SELECT id(),username,password,email,created_at,updated_at FROM users 
	WHERE email==$1 ;
	`
	u := &User{}
	err := store.QueryRow(query, email).Scan(
		&u.ID,
		&u.Name,
		&u.Password,
		&u.Email,
		&u.CreatedAt,
		&u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}
