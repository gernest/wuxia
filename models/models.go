package models

import (
	"time"

	"github.com/gernest/wuxia/db"
)

type Session struct {
	ID        int64
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
	  INSERT INTO sessions VALUES (id(),$1,$2,$3,$4,$5);
	COMMIT
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
	return nil
}
