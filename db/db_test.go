package db

import "testing"

func TestDB(t *testing.T) {
	_, err := Open("ql", "yup.db")
	if err != nil {
		t.Error(err)
	}
}
