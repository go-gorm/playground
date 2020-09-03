package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	city := City{Name: "City-Test"}
	county := County{Name: "County-Test"}

	depository := Depository{
		Name:   "Depository-Test",
		County: county,
		City:   city,
	}

	log.Println("Entity to save: ", depository)

	DB.Create(&depository)

	var depositoryResult Depository
	if err := DB.Set("gorm:auto_preload", true).First(&depositoryResult, depository.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	log.Println("Depository: ", depositoryResult)
	log.Println("Depository City Name: ", depositoryResult.City.Name)

	assert.Equal(t, depository.City, depositoryResult.City)
}
