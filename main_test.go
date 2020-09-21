package main

import (
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	if err := DB.Where(User{Model: gorm.Model{ID: user.ID}}).Update("name", "foo").Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
