package main

import (
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestCreate(t *testing.T) {
	if err := DB.Create(&User{
		Name:        "bob",
		Email:       "bob@gmail.com",
		IntervalStr: "5s",
		Interval:    time.Second * 5,
	}).Error; err != nil {
		t.Fatal(err)
	}
}


func TestUpdate(t *testing.T) {
	if err := DB.Model(User{Id: 1}).Updates(&User{
		Name:        "bob",
		Email:       "bob@gmail.com",
		IntervalStr: "10s",
		Interval:    time.Second * 10,
	}).Error; err != nil {
		t.Fatal(err)
	}
}
