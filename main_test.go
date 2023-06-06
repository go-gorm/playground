package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	db := DB

	user1 := User{Name: "jinzhu"}
	db.Create(&user1)

	user2 := User{Name: "bob", Description: "hello"}
	db.Create(&user2)

	db.Delete(&user2)

	user3 := User{Name: "joe", Description: "hello"}
	db.Create(&user3)

	user4 := User{Name: "peter", Description: "hello"}
	err := db.Create(&user4).Error
	if err == nil {
		t.Error("expected duplicate key error")
	}

}
