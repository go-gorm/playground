package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

// Home is an object with a composite primary key (Street, ApartmentNumber, City, State, Zipcode).
// In the case of an apartment, Steet, City, State, and Zipcode can be the same and the ApartmentNumber is different.
// However, for a house, an apartment number is optional.
type Home struct {
	Street          string `gorm:"primarykey"`
	ApartmentNumber string `gorm:"primarykey"`
	City            string `gorm:"primarykey"`
	State           string `gorm:"primarykey"`
	Zipcode         string `gorm:"primarykey"`
	YearBuilt       string
	MarketStatus    string
}

func TestGORM(t *testing.T) {
	if DB.Migrator().HasTable(&Home{}) {
		err := DB.Migrator().DropTable(&Home{})
		if err != nil {
			t.Errorf("Error while dropping table Home: %v", err)
		}
	}

	err := DB.Migrator().CreateTable(&Home{})
	if err != nil {
		t.Errorf("Error while creating table Home: %v", err)
	}

	home := Home{
		Street:          "1111 Main Street",
		ApartmentNumber: "",
		City:            "Palo Alto",
		State:           "CA",
		Zipcode:         "11111",
		YearBuilt:       "1990",
		MarketStatus:    "Listed For Sale",
	}
	// Insert a new Home record
	if err = DB.Save(&home).Error; err != nil {
		t.Errorf("Error while inserting Home record: %v", err)
	}

	// Update MarketStatus of the same Home record.
	home.MarketStatus = "Sold"
	// Try to save it but this operation fails.
	// Instead of UPDATE, INSERT is attempted.
	if err = DB.Save(&home).Error; err != nil {
		t.Errorf("Error while updating existing Home record: %v", err)
	}
}
