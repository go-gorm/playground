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

// Test inserts and updates the user. Default for column Age is null in the model.
func TestUpdateDefaultValue(t *testing.T) {
	user := User{Name: "fblass"}

	DB.Save(&user)

	var result User
	if err := DB.Where("age is null").First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed insert, got error: %v", err)
	}

	user.Name = "f-blass"
	DB.Save(&user)

	var resultUpdated User
	if err := DB.Where("age is null").First(&resultUpdated, user.ID).Error; err != nil {
		t.Errorf("Failed on update, got error: %v", err)
	}
}
