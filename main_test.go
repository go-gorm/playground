package main

import (
	"gorm.io/hints"
	"testing"
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

func TestSharding(t *testing.T) {
	msg1 := Message{Content: "foo"}
	DB.Create(&msg1)

	var result Message
	query := DB.Model(&result).Where("content = ?", "foo").
		Scan(&result)
	if query.Error != nil {
		t.Errorf("Failed, got error: %v", query.Error)
	}
	if result.Content != "msg" {
		t.Errorf("Failed, got %v", result.Content)
	}
}

func TestShardingAndForceIndex(t *testing.T) {
	msg1 := Message{Content: "foo"}
	DB.Create(&msg1)

	var result Message
	query := DB.Model(&result).Where("content = ?", "foo").
		Clauses(hints.ForceIndex("idx_content")). // <-- This would cause an error
		Scan(&result)
	if query.Error != nil {
		t.Errorf("Failed, got error: %v", query.Error)
	}
	if result.Content != "msg" {
		t.Errorf("Failed, got %v", result.Content)
	}
}
