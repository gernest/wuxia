package models

import (
	"time"

	"github.com/gernest/wuxia/db"
)

//Session stores http session information.
type Session struct {
	Key       string
	Data      []byte
	CreatedOn time.Time
	UpdatedOn time.Time
	ExpiresOn time.Time
}

//Create creates a new session entry into the database
func (s *Session) Create(store *db.DB) error {
	var query = `
	BEGIN TRANSACTION;
	  INSERT INTO sessions VALUES ($1,$2,$3,$4,$5);
	COMMIT;
	`
	tx, err := store.Begin()
	if err != nil {
		return err
	}
	now := time.Now()
	_, err = tx.Exec(query, s.Key, s.Data, now, now, s.ExpiresOn)
	if err != nil {
		return err
	}
	return tx.Commit()
}

// FindByKey  queries the database and returns the Sesion with the given key.The
// result is scanned to the Session object receiver.
//
// It is a good idea to call this on Session object with zero values like this
//  s:&Session{}
//  err:=s.FindByKey(1)
//  if err!=nil{
//    // do something
//  }
//  // now you have s populated with the values returned from the database.
func (s *Session) FindByKey(store *db.DB, key string) error {
	var query = `
	SELECT * from sessions WHERE key LIKE $1 LIMIT 1;
	`
	return store.QueryRow(query, key).Scan(
		&s.Key,
		&s.Data,
		&s.CreatedOn,
		&s.UpdatedOn,
		&s.ExpiresOn,
	)
}

//Count counts the number of rows in the sessions table
func (s *Session) Count(store *db.DB) (int, error) {
	var query = `
	SELECT count() FROM sessions;
	`
	var rst int
	err := store.QueryRow(query).Scan(&rst)
	return rst, err
}

//Update updates a session with a given key with the given data.
func (s *Session) Update(store *db.DB, key string, data []byte) error {
	var query = `
	BEGIN TRANSACTION;
	  UPDATE sessions
		  data = $2,
			updated_on = now(),
    WHERE key==$1;
	COMMIT;
	`
	tx, err := store.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, key, data)
	if err != nil {
		return err
	}
	return tx.Commit()
}

//Delete deletes a row in session table with the given key
func (s *Session) Delete(store *db.DB, key string) error {
	var query = `
	BEGIN TRANSACTION;
	  DELETE FROM sessions
    WHERE key==$1;
	COMMIT;
	`
	tx, err := store.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, key)
	if err != nil {
		return err
	}
	return tx.Commit()
}
