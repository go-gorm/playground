package main

import (
	"database/sql"
	"gorm.io/gorm/clause"
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var result2 User
	stmt := DB.Session(&gorm.Session{DryRun: true}).Where(User{Nickname: sql.NullString{String: "Value", Valid: true}}).First(&result2).Statement
	if len(stmt.Clauses["WHERE"].Expression.(clause.Where).Exprs) < 2 {
		t.Errorf("Failed, WHERE \"users\".\"nickname\" = \"Value\" clause not added.")
	}

	stmt = DB.Session(&gorm.Session{DryRun: true}).Where(User{Nickname: sql.NullString{String: "", Valid: false}}).First(&result2).Statement
	if len(stmt.Clauses["WHERE"].Expression.(clause.Where).Exprs) < 2 { // 2 because it will check WHERE "users"."deleted_at".
		t.Errorf("Failed, WHERE \"users\".\"nickname\" IS NULL clause not added.")
	}
}
