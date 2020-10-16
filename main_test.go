package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	if err := DB.AutoMigrate(&SomeModel{}).Error; err != nil {
		t.Errorf("failed to migrate some model: %v", err)
	}
	if err := DB.AutoMigrate(&SomeAttModel{}).Error; err != nil {
		t.Errorf("failed to migrate some att model: %v", err)
	}
	// user := User{Name: "jinzhu"}

	// DB.Create(&user)

	// var result User
	// if err := DB.First(&result, user.ID).Error; err != nil {
	// 	t.Errorf("Failed, got error: %v", err)
	// }
}
