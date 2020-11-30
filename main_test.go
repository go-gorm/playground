package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	t.Run("20_or_higher_age_validation", func(t *testing.T) {
		DB.Begin()
		defer DB.Rollback()

		user := User{
			Name: "jinzhu",
			Age:  20,
		}
		DB.Create(&user)

		err := DB.Model(&user).Updates(User{
			Age: 19,
		}).Error

		if err == nil {
			t.Error("20 or higher age validation doesn't work")
		}
	})
}
