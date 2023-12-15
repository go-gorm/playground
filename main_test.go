package main

import (
	"gorm.io/gorm/clause"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	c1_1 := Country1{Name: "c1_1"}
	c1_2 := Country1{Name: "c1_2"}
	c2_1 := Country2{CName: "c2_1"}
	c2_2 := Country2{CName: "c2_2"}

	org := Org{
		Adress1_1: Address1{Country: c1_1},
		Adress1_2: Address1{Country: c1_2},
		Adress2_1: Address2{Country: c2_1},
		Adress2_2: Address2{Country: c2_2},
	}
	DB.Create(&org)

	var result Org
	DB.Preload(clause.Associations).First(&result)

	if result.Adress1_1.Country.Name != "c1_1" {
		t.Error("Adress1_1 country has not been resolved")
	}
	if result.Adress1_2.Country.Name != "c1_2" {
		t.Error("Adress1_2 country has not been resolved")
	}
	if result.Adress2_1.Country.CName != "c2_1" {
		t.Error("Adress2_1 country has not been resolved")
	}
	if result.Adress2_2.Country.CName != "c2_2" {
		t.Error("Adress2_2 country has not been resolved")
	}
}
