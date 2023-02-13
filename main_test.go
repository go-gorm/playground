package main

import (
	"testing"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{}

	sql := DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Joins("Company").Take(&user)
	})
	println(sql)
	
	sql = DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Joins("Company", tx.Omit("Name")).Take(&user)
	})
	println(sql)
}
