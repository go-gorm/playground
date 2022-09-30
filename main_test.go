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

func TestBeforeHookWorksWhenObjectIsLoaded(t *testing.T) {
	company := Company{Name: "Gopherz"}

	err := DB.Create(&company).Error
	if err != nil {
		t.Errorf("could not create new company: %v", err)
	}

	if err := DB.First(&Company{}, "id = ?", company.ID).Update("name", "Gophers").Error; err != nil {
		t.Errorf("could not update company name: %v", err)
	}
}

func TestBeforeSaveHookWithoutLoadingObject(t *testing.T) {
	company := Company{Name: "Gopherz"}
	err := DB.Create(&company).Error
	if err != nil {
		t.Errorf("could not create new company: %v", err)
	}

	if err := DB.Model(&Company{}).Where("id = ?", company.ID).Update("name", "Gophers"); err != nil {
		t.Errorf("cannot update name to Gophers! %v", err)
	}
}
