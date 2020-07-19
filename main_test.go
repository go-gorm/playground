package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	var rows []Order

	if err := DB.Model(Order{}).Find(&rows).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Model(Order{}).Where("status = ?", 1).Find(&rows).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Model(&Order{}).Where(&Order{Statue: 1}).Find(&rows);err!= nil{
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Model(Order{}).Table("`order`").Find(&rows).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
