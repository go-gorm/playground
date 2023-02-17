package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Post struct {
	ID uint
	Title string `gorm:"notnull"`
}

func TestGORM(t *testing.T) {
	DB.Debug().AutoMigrate(new(Post))
}
