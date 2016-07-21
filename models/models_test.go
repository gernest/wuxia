package models

import (
	"log"
	"testing"

	"github.com/gernest/wuxia/db"
	"github.com/gernest/wuxia/migration"
)

var store *db.DB

func init() {
	st, err := db.Open("ql-mem", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	store = st
	_ = migration.Up(store)
}

func TestSessions(t *testing.T) {
	sess := &Session{
		Key:  "hello",
		Data: []byte("world"),
	}
	err := sess.Create(store)
	if err != nil {
		t.Error(err)
	}
}
