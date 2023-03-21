package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	test := Test{Whatever: "asdsada"}

	DB.AutoMigrate(&Test{})

	DB.Create(&test)
}
