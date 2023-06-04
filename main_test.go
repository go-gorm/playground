package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres

func TestGORM(t *testing.T) {
	company := Company{
		Name: "Google",
	}

	DB.Create(&company)

	newUser := User{
		Name:      "jinzhu",
		CompanyID: company.ID,
	}

	DB.Create(&newUser)

	var user User
	DB.Preload("Company").First(&user, newUser.ID)

	DB.Model(&user).Updates(User{
		Name: "jinzhu 2",
	})

	var numCompanies int64
	DB.Model(&Company{}).Count(&numCompanies)

	if numCompanies != 1 {
		t.Errorf("There is not 1 company, there are %d", numCompanies)
	}
}
