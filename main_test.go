package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	var userFollows []*UserFollow
	if err := DB.Where("user_id = ?", 79535114761158382).Find(&userFollows).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
