package main

import (
	"gorm.io/gorm/clause"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	DB = DB.Omit(clause.Associations)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	newUser := User{Name: "Testy McBoatFace", Age: 23, Account: Account{Number: "3"}}
	if err := DB.Create(&newUser).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}


	newCompany := Company{Name: "ErrorCo"}
	if err := DB.Create(&newCompany).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}