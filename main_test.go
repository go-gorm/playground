package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	err := DB.Where(&User{Name: "foobar", Age: 32}).Updates(map[string]interface{}{
		"Active": true,
	}).Error
	if  err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	err = DB.Model(&User{Name: "foobar", Age: 32}).Updates(map[string]interface{}{
		"Active": true,
	}).Error
	if  err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
