package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user1 := User{Common: Common{TenantID: 42}, UserID: "one", Name: "jinzhu"}
	user2 := User{Common: Common{TenantID: 42}, UserID: "two", Name: "jinzhu2"}

	DB.Create(&user1)

	if err := DB.Create(&user2).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
