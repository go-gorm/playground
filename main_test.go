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

	vehicle := Vehicle{VehicleID: "id-1"}
	DB.Create(&vehicle)
	var result2 Vehicle
	if err := DB.First(&result2, vehicle.VehicleID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
