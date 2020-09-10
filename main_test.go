package main

import (
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	oldRule := &Rule{
		Name:        "bob",
		Email:       "bob@gmail.com",
		IntervalStr: "5s",
		Interval:    time.Second * 5,
	}
	if err := DB.Create(oldRule).Error; err != nil {
		t.Fatal(err)
	}

	if err := DB.Model(Rule{Id: oldRule.Id}).Updates(&User{
		Name:        "bob",
		Email:       "bob@gmail.com",
		IntervalStr: "10s",
		Interval:    time.Second * 10,
	}).Error; err != nil {
		t.Fatal(err)
	}

	var newRule Rule
	if err := DB.First(&newRule, oldRule.Id).Error; err != nil {
		t.Fatal(err)
	}

	if newRule.IntervalStr != "10s" {
		t.Fatal("wrong field value")
	}

}
