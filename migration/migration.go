package migration

//go:generate go-bindata -o data.go -pkg migration  scripts/...
import "github.com/gernest/wuxia/db"

func Up(store *db.DB) error {
	fname := migrationFile(store.Kind())
	d, err := Asset(fname)
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

func migrationFile(typ db.Kind) string {
	var rst string
	switch typ {
	case db.QL:
		rst = "scripts/up.ql"
	case db.Postgres:
		rst = "scripts/up_pq.sal"
	}
	return rst
}
