package db

import "fmt"

//QLQeryer implements the Queryer interface for ql database.
type QLQeryer struct {
}

//CreateSession retruns a Query for creating new session.
func (ql QLQeryer) CreateSession(table string) Query {
	var query = `
	BEGIN TRANSACTION;
	  INSERT INTO %s (key, data, created_on, updated_on, expires_on)
		VALUES ($1,$2,now(),now(),$3);
	COMMIT;
	`
	query = fmt.Sprintf(query, table)
	return NewQuery(query, true, "key", "data", "expires_on")
}

//FindSessionByKey returns a query for finding a session by key.
func (ql QLQeryer) FindSessionByKey(table string) Query {
	var query = `
	SELECT * from %s WHERE key LIKE $1 LIMIT 1;
	`
	query = fmt.Sprintf(query, table)
	return NewQuery(query, false, "key")
}

//UpdateSession updaates session data.
func (ql QLQeryer) UpdateSession(table string) Query {
	var query = `
BEGIN TRANSACTION;
  UPDATE %s
    data = $2,
    updated_on = now(),
  WHERE key==$1;
COMMIT;
	`
	query = fmt.Sprintf(query, table)
	return NewQuery(query, true, "key", "data")
}

//DeleteSession deletes a session.
func (ql QLQeryer) DeleteSession(table string) Query {
	var query = `
BEGIN TRANSACTION;
   DELETE FROM %s
  WHERE key==$1;
COMMIT;
	`
	query = fmt.Sprintf(query, table)
	return NewQuery(query, true, "key")
}

func (ql QLQeryer) CreateUser(table string) Query {
	var query = `
	BEGIN TRANSACTION;
	  INSERT INTO %s (username,password,email,created_at,updated_at)
		VALUES ($1,$2,$3,now(),now());
	COMMIT;
	`
	query = fmt.Sprintf(query, table)
	return NewQuery(query, true, "username", "password", "email")
}

func (ql QLQeryer) FindUserBy(table, field string) Query {
	var query = `
	SELECT * from %s WHERE %s LIKE $1 LIMIT 1;
	`
	query = fmt.Sprintf(query, table, field)
	return NewQuery(query, false, field)
}
