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

	if err := DB.Model(&user).Updates(&User{Age: 18}).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if user.Age != 18 {
		t.Errorf("Update Age Failed, got : %d , expected : %d", user.Age, 18)
	}
}
