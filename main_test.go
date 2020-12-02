package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu", Age: 30}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var count int64
	var all []User
	if err := DB.Model(new(User)).Select("age").Count(&count).Find(&all).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if all[0].Name == "jinzhu" {
		t.Error("Failed, Select query is age, name should is null")
	}

}
