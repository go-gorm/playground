package main

import (
	"fmt"
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

func TestGORM(t *testing.T) {
	RunMigrations()
	user := User{
		Name: "jinzhu",
		DOB:  time.Now(),
		Age:  0,
	}

	DB.Create(&user)

	DB.Model(&user).Updates(User{
		Name: "het",
		Age:  0,
	}).First(&user)

	fmt.Println(user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
