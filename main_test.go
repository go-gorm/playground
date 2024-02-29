package main

import (
	"gorm.io/gorm"
	"testing"
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

type GlobalOption struct {
	gorm.Model
	Key   string `gorm:"unique;not null;index"`
	Value string
}

func TestBigSerial(t *testing.T) {
	// The old version of gorm.Model (and pg dialect) produced regular integer ID.
	// The current version defaults to `bigserial`, so after upgrade, a migration is attempted.
	// This SQL is directly from pg_dump (minus table namespace):
	err := DB.Exec(`
		CREATE TABLE global_options (
		    id integer NOT NULL,
		    created_at timestamp with time zone,
		    updated_at timestamp with time zone,
		    deleted_at timestamp with time zone,
		    key text NOT NULL,
		    value text
		);
		`).Error
	if err != nil {
		t.Errorf("Initial setup failed, got error: %v", err)
	}
	err = DB.AutoMigrate(&GlobalOption{})
	if err != nil {
		t.Errorf("AutoMigrate failed, got error: %v", err)
	}
}
