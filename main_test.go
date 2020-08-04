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

	var result User
	if DB.First(&result, user.ID); !reflect.DeepEqual(result.CreatedAt, user.CreatedAt) {
		t.Errorf("Failed, result.CreatedAt != user.CreatedAt, result.CreatedAt: %v, user.CreatedAt: %v", result.CreatedAt, user.CreatedAt)
	}
}
