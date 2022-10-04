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

func TestLimitZero(t *testing.T) {
	type LimitZero struct {
		Name string
	}

	DB.Migrator().DropTable(&LimitZero{})

	if err := DB.AutoMigrate(&LimitZero{}); err != nil {
		t.Fatalf("Failed to migrate for uuid default value, got error: %v", err)
	}

	lz := LimitZero{Name: "jinzhu"}
	if err := DB.Create(&lz).Error; err != nil {
		t.Fatalf("should be able to create data, but got %v", err)
	}

	result := make([]LimitZero, 0)
	if err := DB.Model(LimitZero{}).Limit(0).Find(&result).Error; err != nil {
		t.Errorf("No error should happen, but got %v", err)
	}

	if len(result) > 0 {
		t.Fatal("should have zero length")
	}
}
