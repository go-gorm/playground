package main

import (
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestUpdateTime(t *testing.T) {
	prices := []Price{
		{ProductId: 1, Price: 100, SomeTime: time.Now()},
		{ProductId: 2, Price: 150, SomeTime: time.Now()},
	}

	if err := DB.Save(prices).Error; err != nil {
		t.Errorf("Failed to save: %v", err)
	}

	// Change time for the second model, and persist models
	tim := time.Now().Add(-10000 * time.Hour)
	prices[1].SomeTime = tim

	if err := DB.Save(prices).Error; err != nil {
		t.Errorf("Failed to save: %v", err)
	}

	var p Price
	if err := DB.First(&p, 2).Error; err != nil {
		t.Errorf("Failed to get model: %v", err)
	}

	if !p.SomeTime.Equal(tim) {
		t.Errorf("Time was not updated: %v != %v", p.SomeTime, tim)
	}
}
