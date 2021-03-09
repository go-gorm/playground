package main

import (
	"strings"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

func TestGORM(t *testing.T) {

	four := 4

	err := DB.Create(&User{
		Name:   "jinzhu",
		TestID: &four,
	}).Error
	if err == nil {
		t.Fatalf("Expected user creation to fail due to foreign key constraint")
	}
	if !strings.Contains(err.Error(), "violates foreign key constraint") {
		t.Fatalf("Expected foreign key constraint error; got: %s", err)
	}
}
