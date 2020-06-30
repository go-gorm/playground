package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Failed, got error: %v", r)
		}
	}()
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var entity = struct {
		ID int64
	}{}

	if err := DB.Model(&User{}).Select("id").First(&entity, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
