package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

// Reusable scope
func AccountFilter(id string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(&Account{Number: id})
	}
}

func TestGORM(t *testing.T) {
	user := User{
		Name:    "jinzhu",
		Account: Account{Number: "1234"},
	}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	result = User{}
	if err := DB.InnerJoins("Account").First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	result = User{}
	if err := DB.InnerJoins("Account", DB.Where(&Account{Number: "1235"})).First(&result).Error; err == nil {
		t.Errorf("Failed, Return success with invalid account ID")
	}

	result = User{}
	if err := DB.InnerJoins("Account", DB.Scopes(AccountFilter("1235"))).First(&result).Error; err == nil {
		t.Errorf("Failed, Return success with invalid account ID")
	}
}
