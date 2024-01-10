package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	tx := DB.Begin()
	tx.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	tx.Rollback()
}

func TestJoinsWithoutCount(t *testing.T) {
	users := []User{
		{
			Name: "John Doe",
			Company: Company{
				Name: "Company A",
			},
		},
		{
			Name: "Jane Doe",
			Company: Company{
				Name: "Company B",
			},
		},
	}
	tx := DB.Begin()
	tx.Create(&users)

	tx.Model(&User{}).Joins("Company").Order("\"Company\".\"name\" ASC")
	var result []User
	if err := tx.Find(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Failed, expected 2 companies, got %d", len(result))
	}
	if result[0].Company.Name != "Company A" {
		t.Errorf("Failed, expected Company A, got %s", result[0].Company.Name)
	}
	tx.Rollback()
}

func TestJoinsWithCount(t *testing.T) {
	users := []User{
		{
			Name: "John Doe",
			Company: Company{
				Name: "Company A",
			},
		},
		{
			Name: "Jane Doe",
			Company: Company{
				Name: "Company B",
			},
		},
	}
	tx := DB.Begin()
	tx.Create(&users)
	tx.Model(&User{}).Joins("Company").
		Order("\"Company\".\"name\" ASC") // Order still works!

	// --- Critical change
	var count int64
	if err := tx.Count(&count).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	// --- End critical change

	var result []User
	if err := tx.Find(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Failed, expected 2 companies, got %d", len(result))
	}
	if result[0].Company.Name != "Company A" {
		t.Errorf("Failed, expected Company A, got %s", result[0].Company.Name)
	}
	tx.Rollback()
}
