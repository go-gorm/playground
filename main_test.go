package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "test111222333", Age: 20}

	DB.Create(&user)

	var result User
	if err := DB.Omit("age").First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	
	if result.Age == 20 {
		t.Errorf("omit error: %v", result.Age)
	}
}
