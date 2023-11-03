package main

import (
	"errors"
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM1(t *testing.T) {
	user := User{Name: "jinzhu1"}

	var result User
	var errResult error
	db := DB
	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			t.Errorf("failed to create user, err: %s", err.Error())
		}

		// db is parent of this transaction
		errResult = db.First(&result, user.ID).Error

		return nil
	})

	if !errors.Is(errResult, gorm.ErrRecordNotFound) {
		t.Errorf("failed, got error %v", errResult)
	}
}

func TestGORM2(t *testing.T) {
	user := User{Name: "jinzhu2"}

	var result User
	var errResult error
	db := DB.Begin()
	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			t.Errorf("failed to create user, err: %s", err.Error())
		}

		// db is parent of this transaction
		errResult = db.First(&result, user.ID).Error

		return nil
	})

	if !errors.Is(errResult, gorm.ErrRecordNotFound) {
		t.Errorf("failed, got error %v", errResult)
	}
}
