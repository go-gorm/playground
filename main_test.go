package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

type Player struct {
	Name string
	Bio *Bio `gorm:"type:jsonb"`
}

type Bio struct {
	Places []*Places
}

type Places struct {
	Address *string
}

func TestGORM(t *testing.T) {
	player := Player{Name: "jinzhu", Bio: &Bio{}}

	err := DB.Create(&player)
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
