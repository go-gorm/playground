package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu", Languages: []Language{{Name: "EN"}}}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Model(&user).Association("Languages").Replace(Language{Name: "DE"}, Language{Name: "FR"}); err == nil {
		t.Error("Failed, error should be present")
	}
}
