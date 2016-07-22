package models

import (
	"fmt"
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

func CreateSession(store *db.DB, s *Session) error {
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

func FindSessionByKey(store *db.DB, key string) (*Session, error) {
	var query = `
	SELECT * from sessions WHERE key LIKE $1 LIMIT 1;
	`
	s := &Session{}
	err := store.QueryRow(query, key).Scan(
		&s.Key,
		&s.Data,
		&s.CreatedOn,
		&s.UpdatedOn,
		&s.ExpiresOn,
	)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func Count(store *db.DB, table string) (int, error) {
	query := "SELECT count() FROM %s;"
	query = fmt.Sprintf(query, table)
	var rst int
	err := store.QueryRow(query).Scan(&rst)
	return rst, err
}

func UpdateSession(store *db.DB, key string, data []byte) error {
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

func DeleteSession(store *db.DB, key string) error {
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
