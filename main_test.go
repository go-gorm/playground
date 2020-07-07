package main

import (
	"testing"
	
	"github.com/lib/pq"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	message := Message{Title: "", Body: "", Slug: ""}

	err := DB.Create(&message).Error

	if err != nil {
		t.Errorf("Failed, got error: %v", err.(*pq.Error))
	}
}
