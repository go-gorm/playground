package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	// error CREATE INDEX `type` USING TYPE ON `mysq_models`(`name`)
	err := DB.AutoMigrate(&MysqModel{})
	if err != nil {
		t.Fatal(err)
	}
}

type MysqModel struct {
	gorm.Model
	Name string `gorm:"index:type"`
}
