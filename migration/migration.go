//go:generate go-bindata -o data.go -pkg migration  scripts/...
package migration

import "github.com/gernest/wuxia/db"

func Up(store *db.DB) error {
	d, err := Asset("scripts/up.ql")
	if err != nil {
		return err
	}
	tx, err := store.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Commit() }()
	_, err = tx.Exec(string(d))
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
