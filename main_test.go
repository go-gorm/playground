package main

import (
	"os"
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
}

func TestRawWithAtSign(t *testing.T) {
	if os.Getenv("GORM_DIALECT") != "mysql" {
		t.Skip()
	}
	var result map[string]interface{}
	db := DB.Session(&gorm.Session{DryRun: true}).
		Raw("SELECT * FROM users WHERE name = ? OR name = @@version", "dz").Take(&result)
	sql := db.Statement.SQL.String()
	if sql != "SELECT * FROM users WHERE name = 'dz' OR name = @@version" {
		t.Errorf("wrong SQL")
	}
}
