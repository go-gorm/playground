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

	manager := User{Name: "jinzhu"}
	DB.Create(&manager)

	user := User{Name: "jinzhu"}
	user.Account = account
	user.Manager = &manager

	// Causes this:
	// 	INSERT INTO `accounts` VALUES() ON DUPLICATE KEY UPDATE `user_id`=VALUES(`user_id`)
	//  INSERT INTO `users` VALUES() ON DUPLICATE KEY UPDATE `id`=`id`
	// This will create an empty account and user
	DB.Omit("Account.*,Manager.*").Create(&user)

	var accounts []Account
	DB.Find(&accounts)

	var users []User
	DB.Find(&users)

	fmt.Println("LEN USERS: ")
	fmt.Println(len(users))

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

	if len(users) != 2 {
		t.Fail()
	}
}
