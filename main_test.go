package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{
		Model: gorm.Model{ID: 1},
		Name:  "jinzhu",
		Manager: &User{
			Model: gorm.Model{ID: 2},
			Name:  "manager name",
			Age:   42,
		},
	}

	DB.Create(&user)

	var result *User
	if err := DB.Joins("Manager", DB.Session(&gorm.Session{NewDB: true}).Select("id", "name")).First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
		return
	}

	if result.Manager.Age != 0 {
		t.Errorf("Failed, only expected name to be selected, but got age %v", result.Manager.Age)
	}
}
