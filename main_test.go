package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	sql1 := DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Where(
			tx.Where("pizza = ?", "pepperoni").Where(tx.Where("size = ?", "small").Or("size = ?", "medium")),
		).Or(
			tx.Where("pizza = ?", "hawaiian").Where("size = ?", "xlarge"),
		).Find(&User{})
	})

	sql2 := DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		tx = tx.Model(&User{}) // use db.Model!!!
		return tx.Where(
			tx.Where("pizza = ?", "pepperoni").Where(tx.Where("size = ?", "small").Or("size = ?", "medium")),
		).Or(
			tx.Where("pizza = ?", "hawaiian").Where("size = ?", "xlarge"),
		).Find(&User{})
	})

	if sql1 != sql2 {
		t.Errorf("Not equal!\n%s\n%s", sql1, sql2)
	}
}
