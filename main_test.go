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

	// insert a test data, not set time.
	if err := DB.Exec("INSERT INTO `users` (`name`) VALUES ('test')").Error; err != nil {
		t.Errorf("Failed, User insert error: %v ", err)
	}

	// try, query all
	var users []User
	if err := DB.Find(&User{}).Find(&users).Error; err != nil {
		t.Errorf("Failed, User query all error: %v ", err)
	}

}
