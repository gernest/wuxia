package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gernest/wuxia/db"
)

//User represent the application user. It stores important information about the
//user.
//
// The Name, Email and ID fields are all unique. Password is a hashed password,
// hashed using bcrypt algorithm.
type User struct {
	ID        int
	Name      string
	Email     string
	Password  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

//CreateUser creates a new record in the userstable. The Password field is
//hashed before being stored. The Username is sanitized before storing it to the
//database( Don't trust abybody).
func CreateUser(store *db.DB, u *User) error {
	u.Name = Sanitize(u.Name)
	p, err := HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = p
	var query = `
	BEGIN TRANSACTION;
	  INSERT INTO %s (username,password,email,created_at,updated_at)
		VALUES ($1,$2,$3,now(),now());
	COMMIT;
	`
	query = fmt.Sprintf(query, UserTable)
	tx, err := store.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, u.Name, u.Password, u.Email)
	if err != nil {
		return err
	}
	return tx.Commit()
}

//HashPassword hash pass using bcrypt.
func HashPassword(pass []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
}

//VerifyPass verifies that the hash is the hash of pass using bcrypt.
func VerifyPass(hash, pass []byte) error {
	return bcrypt.CompareHashAndPassword(hash, pass)
}

//FindUserByEmail retrieves the user record with the matching email address.
func FindUserByEmail(store *db.DB, email string) (*User, error) {
	var query = `
	SELECT id(),username,password,email,created_at,updated_at FROM %s
	WHERE email==$1 ;
	`
	query = fmt.Sprintf(query, UserTable)
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
