package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	device := Device{Host: "111.111.111.111", Set: "NY14536", CpuUsage: 6.31}

	DB.Create(&device)

	var result []map[string]interface{}
	fields := []string{"host", "cpu_usage", "set"}
	if err := DB.Debug().Table("devices").Select(fields).Find(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
