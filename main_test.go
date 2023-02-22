package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	DB.Create(&User{ID: "jinzhu1", IssuerID: ""})
	
	DB.Create(&User{ID: "jinzhu2", IssuerID: ""})
}
