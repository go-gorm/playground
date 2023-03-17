package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite

type MyTable struct {
	ID    int `gorm:"primaryKey"`
	Field int `gorm:"not null;index:idx_uniq,unique,where:field=0"`
}

func TestGORM(t *testing.T) {
	filename := "test-db.db"
	_ = os.Remove(filename)
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		t.Error(err)
		return
	}

	if err = db.AutoMigrate(&MyTable{}); err != nil {
		t.Error(err)
		return
	}

	// Init
	f := MyTable{ID: 1, Field: 0} // Uniq
	f2 := MyTable{ID: 2, Field: 1}
	f3 := MyTable{ID: 3, Field: 1}
	db.FirstOrCreate(&f)
	db.FirstOrCreate(&f2)
	db.FirstOrCreate(&f3)

	// UNIQUE constraint failed
	if err = db.AutoMigrate(&MyTable{}); err != nil {
		t.Error(err)
		return
	}
}
