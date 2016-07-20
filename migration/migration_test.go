package migration

import (
	"testing"

	"github.com/gernest/wuxia/db"
)

func TestUp(t *testing.T) {
	dbName := "test.db"
	store, err := db.Open("ql-mem", dbName)
	if err != nil {
		t.Fatal(err)
	}
	err = Up(store)
	if err != nil {
		t.Error(err)
	}
}
