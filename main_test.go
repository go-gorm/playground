package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	db := DB
	if err := db.Create(&Pet{FirstName: "Hasso", LastName: "McDog", Tags: []Tag{{Name: "foo"}}}).Error; err != nil {
		t.Fatal(err)
	}

}
