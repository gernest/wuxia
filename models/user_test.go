package models

import "testing"

func TestUser(t *testing.T) {
	pass := "mypass"
	uname := "gernest"
	usr := &User{
		Name:     uname,
		Password: []byte(pass),
	}

	err := CreateUser(store, usr)
	if err != nil {
		t.Error(err)
	}
	err = VerifyPass(usr.Password, []byte(pass))
	if err != nil {
		t.Error(err)
	}
}
