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

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
func TestAutoMigrateDecimal(t *testing.T) {
	err := DB.AutoMigrate(Change{})
	if err != nil {
		t.Fatal(err.Error())
	}
	err = DB.AutoMigrate(Change{})
	if err != nil {
		t.Fatal(err.Error())
	}
}
