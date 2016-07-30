package models

import "testing"

func TestUser(t *testing.T) {
	pass := "mypass"
	uname := "gernest"
	mail := "hello@test.io"
	usr := &User{
		Name:     uname,
		Password: []byte(pass),
		Email:    mail,
	}

	err := CreateUser(store, usr)
	if err != nil {
		t.Error(err)
	}
	err = VerifyPass(usr.Password, []byte(pass))
	if err != nil {
		t.Error(err)
	}

	nusr, err := FindUserByEmail(store, mail)
	if err != nil {
		t.Fatal(err)
	}
	if nusr.Email != mail {
		t.Errorf("expected %s got %s", mail, nusr.Email)
	}
}
