//Package db provides abstractions that makes it easy to interact with the ql
//database.
package db

import (
	"database/sql"

	// load ql drier
	_ "github.com/cznic/ql/driver"
)

//Kind represent the kind of sql database.
type Kind int

// supported databases
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

//Queryer is an interface for alll Queries  executed by the Wuxia application.
//The methods are few to make sure only important queries that are actually
//needed are represented.
type Queryer interface {
	CreateSession(table string) Query
	FindSessionByKey(table string) Query
	UpdateSession(table string) Query
	DeleteSession(table string) Query
}

// implements the Qery interface. It provides an easy way of constructing Query
// compliant structs.
type baseQuery struct {
	s      string
	isTx   bool
	params []string
}

//NewQuery retruns a Query interface satisfying object.
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

//DB extends sql.DB .
type DB struct {
	*sql.DB
	k Kind
	Queryer
}

//Open onens a database connection and returns a *DB instance which is safe for
//concurrent use.
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

//Kind retruns what kind of the database the instance is connected to.
func (db *DB) Kind() Kind {
	return db.k
}
