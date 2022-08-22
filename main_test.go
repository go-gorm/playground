package main

import (
	"fmt"
	"reflect"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "test1", Age: 20}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	updateName, updateAge := "test2", 22
	err := DB.Exec("update users set name = ? and age = ? where id = ?", updateName, updateAge, user.ID).Error
	if err == nil {
		t.Error("expect err: Truncated incorrect DOUBLE value, but got nil")

		var updatedUser User
		if err := DB.First(&updatedUser, user.ID).Error; err != nil {
			t.Errorf("Failed, got new user error: %v", err)
		}
		if reflect.DeepEqual(result, updatedUser) {
			t.Errorf("Failed, update new user failed! old: %v, new: %v", result, updatedUser)
		}
	} else {
		fmt.Printf("got expect error: %v\n", err)
	}
}
