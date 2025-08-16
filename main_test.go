package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	type config struct {
		ID     uint                   `gorm:"primary_key"`
		Config map[string]interface{} `gorm:"column:config;serializer:json"`
	}
	DB.AutoMigrate(&config{})
	DB.Create(&config{Config: map[string]interface{}{"sss": "2#24"}})
	var result map[string]interface{}
	if err := DB.Model(&config{}).First(&result, 1).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
