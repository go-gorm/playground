package main

import (
	"fmt"
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

func TestDoubleStatementExect(t *testing.T) {
	var row struct {
		NextValue uint64
	}
	sequenceName := fmt.Sprintf("MySequence%d", time.Now().Unix())

	if err := DB.Raw(fmt.Sprintf(`
		CREATE SEQUENCE IF NOT EXISTS %s START with 1;
		SELECT last_value + (CASE WHEN is_called THEN 1 ELSE 0 END) as next_value from %s`,
		sequenceName,
		sequenceName)).Scan(&row).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
