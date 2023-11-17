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

	err := DB.Table(User{}.TableName()).Where("name = ?", "jinzhu").Updates(User{Name: "foo"}).Error
	if err != nil {
		panic(err)
	}
}
