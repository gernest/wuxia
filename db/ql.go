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
