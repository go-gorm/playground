package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{
		Name: "jinzhu",
		CreditCard: CreditCard{
			Number:   "123",
			UserName: "jinzhu",
		},
	}

	DB.Create(&user)

	DB.Model(&user).Association("CreditCard").Replace(&CreditCard{
		Number:   "234",
		UserName: "jinzhu",
	})

	var result User
	if err := DB.Preload("CreditCard").First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	fmt.Println(result.CreditCard)

	if result.CreditCard.Number != "234" {
		t.Log("invalid credit card number")
		t.Fail()
	}
}
