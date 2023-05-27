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

func TestSaveVeterinarian(t *testing.T) {
	vet := Planet{
		PlanetData: PlanetData{
			Name:  "Jane",
			Class: "5555555555",
		},
		IsBig: true,
	}

	DB.Create(&vet)

	if vet.ID.String() == "00000000-0000-0000-0000-000000000000" {
		t.Errorf("Failed. Planet Id is empty")
	}
}
