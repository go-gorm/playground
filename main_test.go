package main

import (
	"database/sql"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)
	DB.Create(&Account{
		UserID: sql.NullInt64{Int64: int64(user.ID), Valid: true},
		Number: "123",
	})

	var result struct {
		Something string
		User
	}

	// works
	err := DB.
		Model(User{}).
		//Preload("Account").
		Select("users.*, 'yo' as something").
		First(&result, user.ID).
		Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// panics
	err = DB.
		Model(User{}).
		Preload("Account").
		Select("users.*, 'yo' as something").
		First(&result, user.ID).
		Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
