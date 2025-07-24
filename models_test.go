package main

import (
	"testing"
	"database/sql"
	"fmt"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestUserCreate(t *testing.T) {
	user := User{Name: "Test User No. 2"}

	DB.Create(&user)

	if user.ID != 0 {
		fmt.Printf("User '%s': User (%d) was created\n", user.Name, user.ID)
	} else {
		t.Errorf("User '%s': User creation failed\n", user.Name)
	}

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v\n", err)
	}

	if result.ID != 0 {
		fmt.Printf("User (%d): User (%d) was fetched\n", user.ID, result.ID)
	} else {
		t.Errorf("User (%d): User could not be fetched\n", user.ID)
	}
}

func TestUserAccountWithCompanyCreate(t *testing.T) {
	company := Company{Name: "Test Company No. 3"}

	DB.Create(&company)

	if company.ID != 0 {
		fmt.Printf("Company '%s': Company (%d) was created\n", company.Name, company.ID)
	} else {
		t.Errorf("Company '%s': Company creation failed\n", company.Name)
	}

	user := User{Name: "Test User No. 3", CompanyID: &company.ID, Company: company}

	DB.Create(&user)

	if user.ID != 0 {
		fmt.Printf("User '%s': User (%d) was created\n", user.Name, user.ID)
	} else {
		t.Errorf("User '%s': User creation failed\n", user.Name)
	}

	account := Account{Number: "UserAccount-3", UserID: sql.NullInt64{Int64: int64(user.ID), Valid: true}}

	DB.Create(&account)

	if account.ID != 0 {
		fmt.Printf("Account '%s': Account (%d) was created\n", account.Number, account.ID)
	} else {
		t.Errorf("Account '%s': Account creation failed\n", account.Number)
	}

	var fetchedUser User
	var fetchedAccount Account

	if err := DB.First(&fetchedUser, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v\n", err)
	}

	if fetchedUser.ID != 0 {
		fmt.Printf("User (%d): User (%d) was fetched\n", fetchedUser.ID, fetchedUser.ID)
	} else {
		t.Errorf("User (%d): User could not be fetched\n", user.ID)
	}

	if err := DB.Where("user_id", user.ID).First(&fetchedAccount).Error; err != nil {
		t.Errorf("Failed, got error: %v\n", err)
	}

	if fetchedAccount.ID != 0 {
		fmt.Printf("Account for User (%d): Account Number '%s' was fetched\n", user.ID, fetchedAccount.Number)
	} else {
		t.Errorf("Account for User (%d): Account could not be fetched\n", user.ID)
	}
}
