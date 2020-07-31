package main

import (
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	type Airport struct {
		gorm.Model
		Uuid string `gorm:"type:char(22);not null;uniqueIndex"`
		CompanyId uint32 `gorm:"not null;uniqueIndex:COMPANY_IATA,COMPANY_ICAO"`
		IataAirportCode string `gorm:"type:char(3);uniqueIndex:COMPANY_IATA"`
		IcaoAirportCode string `gorm:"type:char(4);uniqueIndex:COMPANY_ICAO"`
	}

	if err := DB.AutoMigrate(&Airport{}); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	port := &Airport{
		Uuid:            "4MxjS55RAivKy0edVtZurH",
		CompanyId:       1,
		IataAirportCode: "PVG",
		IcaoAirportCode: "ZSPD",
	}

	if err := DB.Create(port).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	port.Uuid = "3pqmAIxhHGb7V0rZ2Egv1X"
	port.CompanyId = 2

	if err := DB.Create(port).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
