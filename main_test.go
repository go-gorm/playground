package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	person := Person{Id: 1231231231231, Name: "test"}
	if err := DB.
		FirstOrCreate(&person, person).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
		return
	}
	if person.Id == 0 {
		t.Errorf("Failed, got error: id=0")
		return
	}

}
