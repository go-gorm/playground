package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{
		Model: gorm.Model{ID: 1},
		Name:  "jinzhu",
		Account: Account{
			// UserID: 1,
			Number: "myAccount",
		},
	}

	// tx := DB
	tx := DB.Begin()
	defer tx.Rollback()
	tx.Create(&user)

	var result User
	if err := tx.Preload("Account").First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	result.Name = "Timothy"                  // This gets saved.
	result.Account.Number = "My new account" // I expect this should be saved as well.

	// I test on Postgres and this fails. The statement I see during save is
	// INSERT INTO "accounts" ("created_at","updated_at","deleted_at","user_id","number","id") VALUES ('2023-03-17 14:06:38.363','2023-03-17 14:06:38.363',NULL,1,'My new account',1) ON CONFLICT ("id") DO UPDATE SET "user_id"="excluded"."user_id" RETURNING "id"
	// And it's because ON CONFLICT DO UPDATE does not update the fields.
	if err := tx.Save(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var result2 User
	if err := tx.Preload("Account").First(&result2, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	assert.Equal(t, "My new account", result2.Account.Number) // fails
}
