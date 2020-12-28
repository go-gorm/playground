package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	company := Company{
		Code: "Code",
		Name: "Something",
	}
	user1 := User{Name: "jinzhu", Company: company}
	DB.Create(&user1)
	user2 := User{Name: "emanuele", Company: company}
	DB.Create(&user2)

	var result1 User
	if err := DB.First(&result1, user1.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var result2 User
	if err := DB.First(&result2, user2.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result1.CompanyID != result2.CompanyID {
		t.Errorf("Expecting same company id")
	}
}
