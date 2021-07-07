package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type ID uint64

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}
	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// the new code

	var id ID
	err := DB.Raw("select id from users where id = ?", user.ID).Scan(&id).Error
	if uint(id) != user.ID && err == nil {
		t.Errorf("Failed, scan fail and error is nil")
	}

}
