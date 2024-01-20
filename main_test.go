package main

import (
	"fmt"
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	data0 := UnionPkModel{Pk1: 1, Pk2: 0}
	data1 := UnionPkModel{Pk1: 1, Pk2: 1}

	DB.Create(&data0)
	DB.Create(&data1)

	var result0, result1 UnionPkModel
	if err := DB.Where("pk1 = ? AND pk2 = ?", data0.Pk1, data0.Pk2).First(&result0).Error; err != nil {
		t.Errorf("Failed to find result0, got error: %v", err)
	}
	if err := DB.Where("pk1 = ? AND pk2 = ?", data1.Pk1, data1.Pk2).First(&result1).Error; err != nil {
		t.Errorf("Failed to find result1, got error: %v", err)
	}

	result0.Data = fmt.Sprintf("%d", time.Now().Unix())
	if err := DB.Save(&result0).Error; err != nil {
		t.Errorf("Failed to save result0, got error: %v", err)
	}
	result1.Data = fmt.Sprintf("%d", time.Now().Unix())
	if err := DB.Save(&result1).Error; err != nil {
		t.Errorf("Failed to save result1, got error: %v", err)
	}
}
