package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	type CasbinRule struct {
		ID    uint   `gorm:"primaryKey;autoIncrement"`
		Ptype string `gorm:"size:100"`
		V0    string `gorm:"size:100"`
		V1    string `gorm:"size:100"`
		V2    string `gorm:"size:100"`
		V3    string `gorm:"size:100"`
		V4    string `gorm:"size:100"`
		V5    string `gorm:"size:100"`
		V6    string `gorm:"size:25"`
		V7    string `gorm:"size:25"`
	}

	// first create table
	if err := DB.AutoMigrate(&CasbinRule{}); err != nil {
		t.Fatal(err)
	}

	// then modify tabel struct
	if err := DB.Exec(fmt.Sprintf("CREATE UNIQUE INDEX %s ON %s (ptype,v0,v1,v2,v3,v4,v5,v6,v7)", "testIndex", "casbin_rules")).Error; err != nil {
		t.Fatal(err)
	}

	// create table again，then it will give an error
	if err := DB.AutoMigrate(&CasbinRule{}); err != nil {
		t.Fatal(err)
	}
}
