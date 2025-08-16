package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type (
	Product struct {
		ProductID  string `gorm:"primaryKey"`
		MaterialID string
		SalesItem  SalesItem `gorm:"foreignKey:MaterialID"`
	}

	SalesItem struct {
		MaterialID string `gorm:"primaryKey"`
	}
)

func TestGORM(t *testing.T) {
	err := DB.AutoMigrate(&Product{}, &SalesItem{})
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	err = DB.Create(&Product{
		ProductID:  "product",
		MaterialID: "1",
		SalesItem: SalesItem{
			MaterialID: "1",
		},
	}).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	err = DB.Create(&SalesItem{
		MaterialID: "2",
	}).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
