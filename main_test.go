package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	now := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(now.UnixNano())), 0)

	one := One{
		ID:   ulid.MustNew(ulid.Timestamp(now), entropy),
		Name: "Asdf",
	}

	DB.Create(&one)

	var result One
	if err := DB.First(&result, one.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
