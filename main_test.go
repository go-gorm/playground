package main

import (
	"testing"

	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	country := Country{Name: "TestCountry"}

	DB.Create(&country)

	user := User{Name: "jinzhu", WorkAddress: Address{CountryID: country.ID}, HomeAddress: Address{CountryID: country.ID}}
	DB.Create(&user)

	var result User
	DB.Preload(clause.Associations).First(&result)

	if result.WorkAddress.Country != country {
		t.Error("WorkAddress country has not been resolved")
	}
	if result.HomeAddress.Country != country {
		t.Error("HomeAddress country has not been resolved")
	}
}
