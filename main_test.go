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
	
	var result2 User
	if err := DB.Table("user").Where("name = ?", "tpp").Find(&result2).Error; err != nil {
		t.Errorf("Failed, User name tpp not found: %v", err)	
	}
}
