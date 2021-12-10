package main

import (
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func (u *User) BeforeCreate(tx *gorm.DB) error {
	myValue, ok := tx.Get("my_value")
	// ok => true
	// myValue => 123
	if ok {
		u.MyValue = myValue.(string)
	}

	return nil
}

func (a *Account) BeforeCreate(tx *gorm.DB) error {
	myValue, ok := tx.Get("my_value")
	// ok => true
	// myValue => 123
	if ok {
		a.MyValue = myValue.(string)
	}

	return nil
}
func TestGORM(t *testing.T) {
	myValue := "bar"
	user := User{Name: "jinzhu", Account: Account{Number: "foo"}}

	DB.Set("my_value", myValue).Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result.MyValue != myValue {
		t.Errorf("Failed, user.MyValue: '%s' want: '%s'", result.MyValue, myValue)
	}

	var accountResult Account
	if err := DB.First(&result, user.Account.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if accountResult.MyValue != myValue {
		t.Errorf("Failed, account.MyValue: '%s' want: '%s'", accountResult.MyValue, myValue)
	}
}
