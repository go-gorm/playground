package main

import (
    "testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Author struct {
    ID      string `gorm:"primaryKey"`
    Books []Book   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;unique"`
}

type Book struct {
    ID       string `gorm:"primaryKey"`
    AuthorID string
}

func TestGORM(t *testing.T) {
    DB.AutoMigrate(&Author{}, &Book{})
}
