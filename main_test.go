package main

import (
	"testing"
	"gorm.io/gorm"
	"fmt"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result map[string]interface{}

	stmt:=DB.Session(&gorm.Session{DryRun: true}).Table("users").First(&result).Statement

	fmt.Println(stmt.SQL.String()) //SELECT * FROM `users` ORDER BY `users`. LIMIT 1

	if err := DB.Table("users").First(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
