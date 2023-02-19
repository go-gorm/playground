package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	userFollow := UserFollow{
		UserID:      79535114761158382,
		FollowedUID: 79542521267813869,
	}

	DB.Create(&userFollow)

	var result UserFollow
	if err := DB.First(&result, userFollow.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
