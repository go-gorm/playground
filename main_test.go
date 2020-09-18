package main

import (
	"errors"
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Asset struct {
	ID         int
	Kind       string
	Value      float32
	BusinessID int
}
type Business struct {
	ID       int
	Name     string
	Assets   []Asset `gorm:"foreignkey:BusinessID;"`
	PersonID int
}
type Person struct {
	ID       int
	Name     string
	Business Business   `gorm:"foreignkey:PersonID;"`
}

func TestGORM(t *testing.T) {

	err := DB.AutoMigrate(&Person{}, &Business{}, &Asset{})
	if err != nil {
		panic(err)
	}
	p := &Person{}
	result := DB.Find(p, "id = ?", "nothing")

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		t.Fail()
	}

}
