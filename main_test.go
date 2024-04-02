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

	userTwo := User{Name: "second"}

	DB.Create(&userTwo)

	var users []User
	if err := DB.Find(&users).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if len(users) != 2 {
		t.Errorf("Expected length 2 go: %v", len(users))
	}

	if users[0].ExternalKey == users[1].ExternalKey {
		t.Errorf("Expected externalKey to be different got: %v", users[0].ExternalKey)
	}
}
