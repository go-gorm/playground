package main

import (
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	company := Company{
		Name: "comp",
	}
	if err := DB.Session(&gorm.Session{NewDB: true}).Save(&company).Error; err != nil {
		t.Errorf("Failed creation, got error: %v", err)
		return
	}
	user := User{
		Name:    "jinzhu",
		Company: company,
	}

	if err := DB.Session(&gorm.Session{NewDB: true}).Save(&user).Error; err != nil {
		t.Errorf("Failed creation, got error: %v", err)
		return
	}

	var result User
	if err := DB.Preload(clause.Associations).First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
		return
	}
	if result.CompanyID == nil {
		t.Errorf("company id is null")
	}

}
