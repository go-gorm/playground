package main

import (
	"errors"
	"testing"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "jinzhu",
	}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
		return
	}

	if err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			if isDuplicate(err) {
				t.Log("--------- DUPLICATE ---------")
				if err := tx.Where("id = ?", user.ID).Delete(&user).Error; err != nil {
					return err
				}
				return nil
			}

			return err
		}

		return nil
	}); err != nil {
		t.Errorf("Failed, got error: %v", err)
		return
	}
}

func isDuplicate(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == pgerrcode.UniqueViolation {
			return true
		}
	}
	return false
}
