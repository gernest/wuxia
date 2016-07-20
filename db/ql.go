//Package db provides abstractions that makes it easy to interact with the ql
//database.
package db

import (
	"database/sql"

	_ "github.com/cznic/ql/driver"
)

type DB struct {
	*sql.DB
}

func Open(dbName, path string) (*DB, error) {
	db, err := sql.Open(dbName, path)
	if err != nil {
		return nil, err
	}
	return &DB{
		DB: db,
	}, nil
}
