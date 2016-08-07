package migration

//go:generate go-bindata -o data.go -pkg migration  scripts/...
import "github.com/gernest/wuxia/db"

type Kind int

const (
	QL Kind = iota
	Postgres
	MySQL
)

func Up(store *db.DB, typ Kind) error {
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

func migrationFile(typ Kind) string {
	var rst string
	switch typ {
	case QL:
		rst = "script/up.ql"
	case Postgres:
		rst = "scripts/up_pq.sal"
	}
	return rst
}
