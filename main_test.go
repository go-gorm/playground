package main

import (
	"time"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql

type User3 struct {
	Id  int64  `gorm:"primaryKey;size:64;autoIncrement:false"`
	Val string `gorm:"size:32"`
	CreatedAt time.Time
}

func TestGORM(t *testing.T) {
	user3 := User3{Id: 1, Val: "jinzhu"}
	DB.AutoMigrate(&User3{})

	if err := DB.Create(&user3).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
