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
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	user.Company.Name = "test"
	err := DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(&user).Error
	if err != nil {
		t.Errorf("Failed upsert user: %+v", err)
		return
	}

	err = DB.First(&result, user.ID).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
		return
	}

	if result.Company.Name != "test" {
		t.Errorf("want test, but got %s", result.Company.Name)
		return
	}

}
