package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	DB.Create(&User{Name: "foo"})
	DB.Create(&User{Name: "jinzhu"})

	var u User
	result := DB.Model(User{
		Name: "jinzhu",
	}).First(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	fmt.Println(u.Name)
	if u.Name != "jinzhu" {
		t.Errorf("Failed, got record: %s", u.Name)
	}
}
