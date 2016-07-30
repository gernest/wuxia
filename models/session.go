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

//CreateSession creates a new database record for the session s object. This
//sets CreatedOn and UpdatedOn fieds to time.Now().
func CreateSession(store *db.DB, s *Session) error {
	var query = `
	BEGIN TRANSACTION;
	  INSERT INTO %s VALUES ($1,$2,$3,$4,$5);
	COMMIT;
	`
	query = fmt.Sprintf(query, SessionTable)
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

//FindSessionByKey queries the databse for session with key field key.
func FindSessionByKey(store *db.DB, key string) (*Session, error) {
	var query = `
	SELECT * from %s WHERE key LIKE $1 LIMIT 1;
	`
	query = fmt.Sprintf(query, SessionTable)
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

//Count counts the number of rows in a table named table.
func Count(store *db.DB, table string) (int, error) {
	query := "SELECT count() FROM %s;"
	query = fmt.Sprintf(query, table)
	var rst int
	err := store.QueryRow(query).Scan(&rst)
	return rst, err
}

//UpdateSession updates a ratabase record for session whose key matches key.
//Only two fields are updated, the data and updated_on field.
func UpdateSession(store *db.DB, key string, data []byte) error {
	var query = `
BEGIN TRANSACTION;
  UPDATE %s
    data = $2,
    updated_on = now(),
  WHERE key==$1;
COMMIT;
	`
	query = fmt.Sprintf(query, SessionTable)
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

//DeleteSession deletes a session record which matches key.
func DeleteSession(store *db.DB, key string) error {
	var query = `
BEGIN TRANSACTION;
   DELETE FROM %s
  WHERE key==$1;
COMMIT;
	`
	query = fmt.Sprintf(query, SessionTable)
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
