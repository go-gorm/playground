package main

import (
	"testing"
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

func TestWhere1(t *testing.T) {
	var result []map[string]interface{}
	if err := DB.Table("users").Where("1").Scan(&result).Error; err != nil {
		t.Errorf("error on selecting * from users where 1: %v", err)
	}
}
