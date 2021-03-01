package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{
		Name: "jinzhu",
		Age:  32,
		Account: Account{
			Number: "123",
		},
	}

	DB.Session(&gorm.Session{NewDB: true}).Create(&user)

	user.Account.Number = "324"
	DB.Session(&gorm.Session{NewDB: true}).Model(User{}).Omit("Account").Save(&user)

	var createdUser User
	if err := DB.Session(&gorm.Session{NewDB: true}).Model(User{}).Find(&createdUser, "id = ?", user.ID).Error; err != nil {
		t.Errorf("there should not be any error finding: %v", err)
	}
	if createdUser.Account.Number != "123" {
		t.Errorf("The number should not have updated: %v", createdUser.Account.Number)
	}
	return

}
