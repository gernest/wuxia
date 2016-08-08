package db

import "fmt"

type QLQeryer struct {
}

func (ql QLQeryer) CreateSession(table string) Query {
	var query = `
	BEGIN TRANSACTION;
	  INSERT INTO %s (key, data, created_on, updated_on, expires_on)
		VALUES ($1,$2,now(),now(),$3);
	COMMIT;
	`
	query = fmt.Sprintf(query, table)
	return NewQuery(query, true, "key", "data", "expire_on")
}
