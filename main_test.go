package main

import (
	"testing"
	"strconv"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var deletedUser User
	err := dbPG.
		Where("created_at < NOW() - INTERVAL ?", strconv.Itoa(10)+" hours").
		Delete(&deletedUser).Error
	
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
