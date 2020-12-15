package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	NormalUpdate(t)
	FailingUpdate(t)

}

func NormalUpdate(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	updates := map[string]interface{}{}
	updates["Name"] = "Blah"

	if err := DB.Model(&result).Updates(updates).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result.Name != "Blah" {
		t.Errorf("Failed, DeletedAt is empty")
	}

	if result.Model.UpdatedAt == user.Model.UpdatedAt {
		t.Errorf("Failed, UpdatedAt not updated")
	}
}

func FailingUpdate(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	updates := map[string]interface{}{}
	updates["Name"] = "Blah"
	fields := []string{"Name"}

	if err := DB.Model(&result).Select(fields).Updates(updates).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result.Name != "Blah" {
		t.Errorf("Failed, DeletedAt is empty")
	}

	if result.Model.UpdatedAt == user.Model.UpdatedAt {
		t.Errorf("Failed, UpdatedAt not updated")
	}

}
