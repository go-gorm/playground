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

	user.Age = 12
	DB.Create(&user)

	err := DB.Model(&user).Updates(User{
		Name: "het",
		Age:  0,
	}).First(&user).Error

	if err != nil {
		fmt.Println("Got Error:", err)
	}
	fmt.Println(user)
}
