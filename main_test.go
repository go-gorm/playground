package main

import (
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver, cockroach

func TestGORM(t *testing.T) {
	// time.Local = time.FixedZone("CET", 3600)
	// time.Local = time.FixedZone("UTC", 0)

	user := User{Name: "jinzhu"}

	DB.Create(&user)
	now := time.Now()

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	createdAtZone, CreatedAtOffset := result.CreatedAt.Zone()
	nowZone, nowOffset := now.Zone()
	t.Logf("NOW: %v, Created At: %v", now, result.CreatedAt)
	t.Logf("Created At Zone name: %s, Created At Offset: %d", createdAtZone, CreatedAtOffset)
	t.Logf("NOW Zone name: %s, NOW Offset: %d", nowZone, nowOffset)

	timeNowString := now.Format("2006-01-02 15:04")
	createdAtString := result.CreatedAt.Format("2006-01-02 15:04")

	t.Logf("NOW: %s, Created At: %s", timeNowString, createdAtString)

	if timeNowString == createdAtString && CreatedAtOffset != nowOffset {
		t.Errorf("Failed, times cannot be the same with different time zones")
	}
}
