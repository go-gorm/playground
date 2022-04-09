package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var u User

	err := DB.
		Select("max(id)").
		Where("date(created_at) >= date(now()) - ?", 0).
		Group("date(created_at)").
		Find(&u)

	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
