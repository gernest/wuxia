package models

import (
	"fmt"
	"time"

	"github.com/gernest/wuxia/db"
)

//Session stores http session information.
type Session struct {
	Key       string `store:"key"`
	Data      []byte `store:"data"`
	CreatedOn time.Time
	UpdatedOn time.Time
	ExpiresOn time.Time `store:"expires_on"`
}

//CreateSession creates a new database record for the session s object. This
//sets CreatedOn and UpdatedOn fieds to time.Now().
func CreateSession(store *db.DB, s *Session) error {
	q := store.CreateSession(SessionTable)
	_, err := ExecModel(store, s, q)
	if err != nil {
		return err
	}
	return nil
}

//FindSessionByKey queries the databse for session with key field key.
func FindSessionByKey(store *db.DB, s *Session) (*Session, error) {
	q := store.FindSessionByKey(SessionTable)
	err := QueryRowModel(store, s, q).Scan(
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
func UpdateSession(store *db.DB, s *Session) error {
	q := store.UpdateSession(SessionTable)
	_, err := ExecModel(store, s, q)
	if err != nil {
		return err
	}
	return nil
}

//DeleteSession deletes a session record which matches key.
func DeleteSession(store *db.DB, s *Session) error {
	q := store.DeleteSession(SessionTable)
	_, err := ExecModel(store, s, q)
	if err != nil {
		return err
	}
	return nil
}
