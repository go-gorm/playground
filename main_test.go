package main

import (
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestPreloadHooks(t *testing.T) {
	u := User{Name: "caleb", Company: Company{Name: "company"}}
	DB.Create(&u)

	// Preload should be skipped
	DB.Session(&gorm.Session{SkipHooks: true}).Model(&User{}).Preload(clause.Associations).First(&u, u.ID)
	if u.Name != "caleb" {
		t.Errorf("want caleb, Failed, got: %v", u.Name)
	}

	if u.Company.Name != "company" {
		t.Errorf("want company, got: %v", u.Company.Name)
	}
}
