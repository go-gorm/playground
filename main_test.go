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

	pet := Pet{UserID: &user.ID, Name: "fido"}

	DB.Create(&pet)

	sqlDB, err := DB.DB()
	if err != nil {
		t.Fatal(err)
	}
	sqlDB.SetMaxOpenConns(1)

	// Rerun the migrations now that data is present.
	if err := DB.AutoMigrate(&User{}, &Pet{}); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
