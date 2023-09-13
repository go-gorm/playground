package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)
	res := DB.Create(&user)
	if res.Error != gorm.ErrDuplicatedKey {
		t.Errorf("Error is not gorm.ErrDuplicateKey")
	}
}
