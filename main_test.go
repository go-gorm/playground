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
		Account: Account{
			Number: "123456",
			Companies: []Company{
				{Name: "Corp1"}, {Name: "Corp2"},
			},
			Pet: Pet{
				Name: "Pet1",
			},
		},
	}

	DB.Create(&user)
	fmt.Println("-------------------------------------------------------")
	var count int64
	var result User
	DB = DB.Model(&User{}).
		Joins("Account").
		Joins("Account.Pet").
		Preload("Account.Companies")

	if err := DB.Count(&count).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if len(result.Account.Companies) != 2 {
		t.Errorf("Failed, got %v", len(result.Account.Companies))
	}

	if result.Account.Pet.Name != "Pet1" {
		t.Errorf("Failed, got '%v'", result.Account.Pet.Name)
	}
}
