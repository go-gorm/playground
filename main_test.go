package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu",Age: 1}
	user1 := User{Name: "dd",Age: 2}

	DB.Create(&user)
	DB.Create(&user1)

	if err := DB.Exec("SELECT count(*) FROM users WHERE Age IN (?)", []int{1,2}).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
