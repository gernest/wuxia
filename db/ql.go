//Package db provides abstractions that makes it easy to interact with the ql
//database.
package db

import (
	"database/sql"

	_ "github.com/cznic/ql/driver"
)

type Kind int

const (
	QL Kind = iota
	Postgres
	MySQL
)

type DB struct {
	*sql.DB
	k Kind
}

func Open(dbName, path string) (*DB, error) {
	db, err := sql.Open(dbName, path)
	if err != nil {
		return nil, err
	}
	var k Kind
	switch dbName {
	case "q;", "ql-mem":
		k = QL
	}
	return &DB{
		DB: db,
		k:  k,
	}, nil
}

func (db *DB) Kind() Kind {
	return db.k
}
