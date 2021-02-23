package main

import (
	"database/sql"
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{
		Name: "jinzhu",
		Account: Account{
			UserID: sql.NullInt64{Int64: 123, Valid: true},
			Number: "123",
		},
	}

	DB.Create(&user)

	var result User
	if err := DB.Preload("Account").First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	oldUpdateAt := result.Account.UpdatedAt
	time.Sleep(1 * time.Second)

	result.Account.Number = "456"
	DB.Save(&result)

	var result2 User
	if err := DB.Preload("Account").First(&result2, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result2.Account.UpdatedAt.Equal(oldUpdateAt) {
		t.Errorf("Failed, UpdatedAt on account not updated, old: %v new: %v", oldUpdateAt, result2.Account.UpdatedAt)
	}

}
