package main

import (	
	"testing"
	"fmt"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	if user.ID != 0 {
		fmt.Printf("User '%s': User (%d) was created\n", user.Name, user.ID)
	} else {
		t.Errorf("User '%s': User creation failed\n", user.Name)
	}

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v\n", err)
	}

	if result.ID != 0 {
		fmt.Printf("User (%d): User (%d) was fetched\n", user.ID, result.ID)
	} else {
		t.Errorf("User (%d): User could not be fetched\n", user.ID)
	}
}
