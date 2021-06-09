package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	// When we mistakenly pass an entire struct, which has a map as one of its fields, as the Scan destination, Scan
	// logs an error, but does not return it.
	//
	// [error] unsupported data type: &map[]

	s := struct {
		I int64
		M map[string]string
	}{}

	err := DB.Raw("SELECT 7").Scan(&s.I).Error
	if err != nil {
		t.Fatal("expected no error")
	}
	if s.I != 7 {
		t.Fatalf("expected s.I to be 7, got %d", s.I)
	}

	// We mistakenly passed "&s" instead of "&s.I"
	err = DB.Raw("SELECT 13").Scan(&s).Error
	if err == nil {
		t.Error("expected an error, got nil")
	}
	if s.I != 7 {
		t.Errorf("expected s.I to be 7, got %d", s.I)
	}
}
