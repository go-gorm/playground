package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	ip := Ip{IP: "1.1.1.1", Flow: 12.3}
	if err := DB.AutoMigrate(&ip); err != nil {
		t.Fatal(err)
	}
	DB.Table(`1.1.1.1`).Create(&ip)

	var count int64
	DB.Table(`1.1.1.1`).Count(&count)
	if count != 1 {
		t.Fatalf("count should be %d", 1)
	}
}
