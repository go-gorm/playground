package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	DB.AutoMigrate(&Person{})
	if err := DB.AutoMigrate(&Cloth{}).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

}
