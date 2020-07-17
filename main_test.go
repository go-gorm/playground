package main

import (
	"testing"
	
	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.Model(&user).Omit(clause.Associations).Updates(map[string]interface{}{"name": "test"}).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
