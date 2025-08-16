package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Part struct {
	PartID  int    `gorm:"column:part_id;not null"`
	Variant string `gorm:"not null"`
}

type Batch struct {
	Part
	ProdCode string `gorm:"primarykey;not null"`
	ProdDate string `gorm:"not null"`
}

type PartCatalog struct {
	IndustryCode string `gorm:"not null"`
	IndustryName string `gorm:"not null"`
	InternalPart Part   `gorm:"embedded;primarykey;not null"`
}

func TestGORM(t *testing.T) {
	DB.AutoMigrate(&PartCatalog{}, &Batch{})

	part1 := &PartCatalog{
		IndustryCode: "RE189",
		IndustryName: "Rotary Expulsion System",
		InternalPart: Part{
			PartID: 19,
			Variant: "3-WH46"
		},
	}

	part2 := &PartCatalog{
		IndustryCode: "AC177",
		IndustryName: "Titanium Ball Bearing",
		InternalPart: Part{
			PartID: 120,
			Variant: "1-BE47"
		},
	}

	DB.Create(&part1, &part2)

	first := &PartCatalog{}
	if err := DB.First(&first); err != nil {
		t.Errorf("Failed, got error: %v", err)
	} else if first.InternalPart.PartID != 19 {
		t.Errorf("Failed, this should not be first by primary key: %v", first.InternalPart.PartID)
	}
}
