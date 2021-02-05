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
}

func TestUserCompany(t *testing.T) {
	user := User{
		Name: "jinzhu",
		Company: Company{
			Name: "TestCompany",
		},
	}

	if err := DB.Create(&user).Error; err != nil {
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
	if result.Company.Name == "" {
		t.Errorf("company not saved")
		return
	}

	user.Company.Name = "NewTestCompany"
	if err := DB.Save(&user).Error; err != nil {
		t.Errorf("Failed to update user. %v", err.Error())
		return
	}

	if err := DB.Preload(clause.Associations).First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
		return
	}
	if result.Company.Name == "NewTestCompany" {
		t.Errorf("company name got updated eventhough it is create only permission")
		return
	}

}
