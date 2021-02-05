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

	if err := DB.Preload(clause.Associations).Where("name = ?", user.Name).Find(&user).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestFirstOrCreate(t *testing.T) {
	user := User{
		Name: "jinzhu",
	}

	if err := DB.Preload(clause.Associations).Where("name = ?", user.Name).Find(&user).Error; err != nil {
		t.Errorf("Failed to create user. %v", err.Error())
		return
	}

	if user.CompanyID == nil || *user.CompanyID == 0 {
		t.Errorf("CompanyID field not updated")
		return
	}
	var result User
	if err := DB.Preload(clause.Associations).First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
		return
	}

}
