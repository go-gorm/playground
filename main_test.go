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

	var intf interface{} = user
	if err := DB.Save(&intf).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
