package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func CreateObject() error {
	type Object struct{}
	return DB.AutoMigrate(&Object{})
}

func AddField1() error {
	type Object struct {
		Field1 string
	}
	return DB.AutoMigrate(&Object{})
}

func AddField2() error {
	type Object struct {
		Field2 string
	}
	return DB.AutoMigrate(&Object{})
}

func TestGORM(t *testing.T) {
	if DB.Dialector.Name() != "postgres" {
		return
	}
	if err := CreateObject(); err != nil {
		t.Errorf("FAILED CREATING OBJECT")
		return
	}
	if err := AddField1(); err != nil {
		t.Errorf("FAILED ADDING FIELD1")
		return
	}
	if err := AddField2(); err != nil {
		t.Errorf("I am sad :(")
		return
	}
}
