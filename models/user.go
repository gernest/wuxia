package models

import (
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
	ID        int64     `store:"id"`
	Name      string    `store:"username"`
	Email     string    `store:"email"`
	Password  []byte    `store:"password"`
	CreatedAt time.Time `store:"created_on"`
	UpdatedAt time.Time `store:"updated_on"`
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
	q := store.CreateUser(UserTable)
	_, err = ExecModel(store, u, q)
	if err != nil {
		return err
	}
	return nil
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
	q := store.FindUserBy(UserTable, "email")
	u := &User{
		Email: email,
	}
	err := QueryRowModel(store, u, q).Scan(
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
