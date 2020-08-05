package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	for i := uint(0); i < 10000; i ++ {
		err := DB.Updates(&user).Where("age = ?", i).Error
		if err != nil {
				t.Errorf("gorm errored :%v ", err)
		}
	}
}