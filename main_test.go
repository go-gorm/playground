package main

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	if err := DB.Transaction(func(tx *gorm.DB) error {
		return tx.AutoMigrate(&User{})
	}); err != nil {
		t.Error(err)
	}

	fmt.Println("AutoMigrate(PostgreSQL,timestamptz(0)) First OK, but when AutoMigrate Again, while be deadlock.")

	if err := DB.Transaction(func(tx *gorm.DB) error {
		return tx.AutoMigrate(&User{})
	}); err != nil {
		t.Error(err)
	}

	fmt.Println("Never go to here ...")
}
