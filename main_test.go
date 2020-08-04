package main

import (
	"testing"
	"reflect"
	"encoding/json"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	DB.First(&result, user.ID)
	
	if !result.CreatedAt.Equal(user.CreatedAt) {
		t.Errorf("time.Equal Failed, result.CreatedAt: %v, user.CreatedAt: %v", result.CreatedAt, user.CreatedAt)
	}
}
