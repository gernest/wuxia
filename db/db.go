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

//Query is an interface for SQL query string.
type Query interface {
	String() string
	IsTx() bool
	Params() []string
}

type Queryer interface {
	CreateSession(table string) Query
}

type baseQuery struct {
	s      string
	isTx   bool
	params []string
}

func NewQuery(q string, tx bool, p ...string) Query {
	return &baseQuery{s: q, isTx: tx, params: p}
}

func (b *baseQuery) String() string {
	return b.s
}

func (b *baseQuery) IsTx() bool {
	return b.isTx
}

func (b *baseQuery) Params() []string {
	return b.params
}

type DB struct {
	*sql.DB
	k Kind
	Queryer
}

func Open(dbName, path string) (*DB, error) {
	db, err := sql.Open(dbName, path)
	if err != nil {
		return nil, err
	}
	var k Kind
	var q Queryer
	switch dbName {
	case "q;", "ql-mem":
		k = QL
		q = QLQeryer{}
	}
	return &DB{
		DB:      db,
		k:       k,
		Queryer: q,
	}, nil
}

func (db *DB) Kind() Kind {
	return db.k
}
