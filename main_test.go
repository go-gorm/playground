package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	name := "name"
	user := map[string]interface{}{name: "id"}
	//user := User{Name: "jinzhu"}

	err := DB.Table("users").Create(user).Error

	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
