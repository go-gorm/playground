package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu", Age: 5}
	toy := Toy{Name: "jinzhu", OwnerID: "test"}

	DB.Create(&user)
	DB.Create(&toy)

	var result User
	DB = DB.Where("name = ?", "jinzhu")

	if err := DB.First(&result, "age = ?", 5).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	var ttoy Toy

	if err := DB.First(&ttoy, "owner_id = ?", "test").Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
