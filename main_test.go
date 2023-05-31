package main

import (
	"testing"

	"github.com/google/uuid"
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

func TestSavePlanet(t *testing.T) {

	t.Run("should get id from db", func(t *testing.T) {
		planet := Planet{
			PlanetData: PlanetData{
				Name:  "Jane",
				Class: "5555555555",
			},
			IsBig: true,
		}

		DB.Create(&planet)

		if planet.ID.String() == "00000000-0000-0000-0000-000000000000" {
			t.Errorf("Failed. Planet Id is empty")
		}

	})

	t.Run("should save planet with given id", func(t *testing.T) {
		expectedId := uuid.New()
		planet := PlanetWithDefault{
			ID: expectedId,
			PlanetData: PlanetData{
				Name:  "Jane",
				Class: "5555555555",
			},
			IsBig: true,
		}

		DB.Create(&planet)

		if planet.ID != expectedId {
			t.Errorf("Failed. Id is not the expected")
		}
	})

}
