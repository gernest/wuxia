package models

import (
	"log"
	"testing"
	"time"

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
		Key:       "hello",
		Data:      []byte("world"),
		ExpiresOn: time.Now(),
	}
	err := CreateSession(store, sess)
	if err != nil {
		t.Error(err)
	}
	n, err := Count(store, "sessions")
	if err != nil {
		t.Error(err)
	}
	if n != 1 {
		t.Errorf("expected 1 got %d", n)
	}

	ns, err := FindSessionByKey(store, &Session{Key: sess.Key})
	if err != nil {
		t.Error(err)
	}
	if ns.Key != sess.Key {
	}
}
