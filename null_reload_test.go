package main

import (
	"testing"
	"time"
)

func TestReloadNullPointer(t *testing.T) {
	if DB.Dialector.Name() != "postgres" {
		t.Skip("This reproduction uses a PostgreSQL temporary table")
	}

	tx := DB.Begin()
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}
	defer tx.Rollback()

	if err := tx.Exec("CREATE TEMP TABLE gorm_null_reload_repro (id bigint PRIMARY KEY, archived_at timestamptz) ON COMMIT DROP").Error; err != nil {
		t.Fatal(err)
	}
	type Record struct {
		ID         int64
		ArchivedAt *time.Time
	}
	timestamp := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	if err := tx.Table("gorm_null_reload_repro").Create(&Record{ID: 1, ArchivedAt: &timestamp}).Error; err != nil {
		t.Fatal(err)
	}

	var record Record
	if err := tx.Table("gorm_null_reload_repro").First(&record, 1).Error; err != nil {
		t.Fatal(err)
	}
	if record.ArchivedAt == nil {
		t.Fatal("initial load unexpectedly NULL")
	}
	if err := tx.Exec("UPDATE gorm_null_reload_repro SET archived_at = NULL WHERE id = 1").Error; err != nil {
		t.Fatal(err)
	}
	if err := tx.Table("gorm_null_reload_repro").First(&record, 1).Error; err != nil {
		t.Fatal(err)
	}

	var fresh Record
	if err := tx.Table("gorm_null_reload_repro").First(&fresh, 1).Error; err != nil {
		t.Fatal(err)
	}
	if fresh.ArchivedAt != nil {
		t.Fatalf("fresh load = %v, expected NULL", fresh.ArchivedAt)
	}
	if record.ArchivedAt != nil {
		t.Fatalf("reused destination retained %v; fresh destination correctly loaded nil", *record.ArchivedAt)
	}
}
