package main

import (
	"gorm.io/gorm"
	"log"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	type Print struct {
		gorm.Model
		OwnerID   uint   `gorm:"not null;index"`
		OwnerType string `gorm:"type:varchar(32);not null;index"`
	}

	DB.Migrator().DropTable(&Print{})
	DB.AutoMigrate(&Print{})

	printIDs := make([]uint, 0)
	ids := []uint{10}

	// Test without parentheses
	err := DB.Model(&Print{}).
		Where("owner_type = ? AND owner_id IN ?", "cart_products", ids).
		Pluck("id", &printIDs).Error

	// Test with parentheses
	err = DB.Model(&Print{}).
		Where("owner_type = ? AND owner_id IN (?)", "cart_products", ids).
		Pluck("id", &printIDs).Error

	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	log.Println(printIDs)
}
