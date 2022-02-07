package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: v1.22.5
// TEST_DRIVERS: mysql

func TestGORM(t *testing.T) {
	account := Account{Number: "1"}
	DB.Create(&account)

	user := User{Name: "jinzhu"}
	user.Account = account

	// Causes this: INSERT INTO `accounts` VALUES() ON DUPLICATE KEY UPDATE `user_id`=VALUES(`user_id`)
	// This will create an empty account
	DB.Omit("Account.*").Create(&user)

	var accounts []Account
	DB.Find(&accounts)

	if len(accounts) != 1 {
		fmt.Print("Num Accounts")
		fmt.Println(len(accounts))

		for _, a := range accounts {
			fmt.Print("Account number: ") // one of these will be an empty string
			fmt.Println(a.Number)
		}
		t.Error("unexpected accounts")
		t.Fail()
	}
}
