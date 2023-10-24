package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	pets := []*Pet{}

	stmt := DB.Session(&gorm.Session{DryRun: true}).
		Select("*").
		Joins("join users on users.id = pets.user_id").
		Where(&User{Name: "user_name"}).Find(&pets).Statement
	expect := "SELECT * FROM `pets` join users on users.id = pets.user_id WHERE `users`.`name` = ? AND `pets`.`deleted_at` IS NULL"

	if got := stmt.SQL.String(); got != expect {
		t.Errorf("\nexpect: %s\n   got: %v", expect, got)
	}

}
