package main

import "testing"

func TestInitUser(t *testing.T) {
	err := initUser()
	if err != nil {
		t.Errorf("Init User failed: %s.", err.Error())
	}
}

func TestInitGit(t *testing.T) {
	TestInitUser(t)

	err := initGit()
	if err != nil {
		t.Errorf("Init Git failed: %s.", err.Error())
	}
}

func TestInitBadger(t *testing.T) {
	TestInitUser(t)

	db, err := initBadger()
	if err != nil {
		t.Errorf("Init Badger failed: %s.", err.Error())
	}

	err = db.Close()
	if err != nil {
		t.Errorf("Close Badger failed: %s.", err.Error())
	}
}
