package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{
		Name: "jinzhu",
	}

	// Check if username is taken
	// Bug description: If we test if the username is taken the Create statement will fail
	// with an error 'record not found'. If we skip this (comment it out) we can
	// successfully create the new user in line 25
	if usernameTaken(user.Name) {
		t.Errorln("Failed, username taken")
	}

	// If not taken, create new user
	// Bug description: This will fail with the above test (usernameTaken) in place
	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func usernameTaken(username string) bool {
	user := new(User)
	DB.Where("name = ?", username).First(&user)
	return user.ID != 0
}
