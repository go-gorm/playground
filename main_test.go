package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Bar struct {
	ID uint `gorm:"primarykey"`
}

type Foo struct {
	ID    uint `gorm:"primarykey"`
	BarID uint `gorm:"index;NOT NULL"`
	Bar   Bar
}

func TestGORM(t *testing.T) {
	DB.Table("foo").AutoMigrate(&Foo{})
}
