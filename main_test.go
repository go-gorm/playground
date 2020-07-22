package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type User2 struct {
	Aid int64  `gorm:"primaryKey;size:64;autoIncrement:false`
	Val string `gorm:"size:32"`
}

func TestGORM(t *testing.T) {
	user2 := User2{Aid: 1, Val: "jinzhu"}

	if err := DB.Create(&user2).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
