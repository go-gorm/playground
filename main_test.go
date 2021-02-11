package main

import (
	"testing"

	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	if err := DB.Preload(clause.Associations).Create(&user).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestArrayCreation(t *testing.T) {
	pets := []*Pet{
		&Pet{
			Name: "Pet1",
		},
		&Pet{
			Name: "Pet2",
		},
	}
	user := User{
		Name: "jinzhu",
		Pets: pets,
	}

	if err := DB.Preload(clause.Associations).Create(&user).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if user.ID == 0 {
		t.Errorf("user should have an id")
	}
	if user.Name != "jinzhu" {
		t.Errorf("user name should be jinzhu")
	}
	for _, pet := range user.Pets {
		if pet.ID == 0 {
			t.Errorf("pet should have an ID after creation")
		}
	}

	user.Name = "abel"
	if err := DB.Preload(clause.Associations).Save(&user).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

}
