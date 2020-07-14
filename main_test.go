package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func fork(tx *gorm.DB) *gorm.DB {
	return tx // This mock function should be replace by `gorm.DB.Fork()` and should implement the fork feature
}

func TestGORM(t *testing.T) {
	DB = DB.Session(&gorm.Session{DryRun: true})

	DB.Create(&User{Name: "User 1", Age: 1})
	DB.Create(&User{Name: "User 2", Age: 10})
	DB.Create(&User{Name: "User 3", Age: 100})
	DB.Create(&User{Name: "User 4", Age: 1000})

	ctx1 := context.Background()

	baseTx := fork(DB).WithContext(ctx1)

	greaterThan10Tx := fork(baseTx).Where("age > 10")

	var greaterThan10 []User
	tx := greaterThan10Tx.Find(&greaterThan10) // Run that with context 1 & age > 10
	assert.Equal(t, "SELECT `id`,`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`birthday`,`company_id`,`manager_id`,`active` FROM `users` WHERE age > 10 AND `users`.`deleted_at` IS NULL", tx.Statement.SQL.String())

	greaterThan10LowerThan1000Tx := fork(greaterThan10Tx).Where("age < 1000")

	var greaterThan10LowerThan1000 []User
	tx = greaterThan10LowerThan1000Tx.Find(&greaterThan10LowerThan1000) // Run that with context 1 & age > 10 & age < 1000
	assert.Equal(t, "SELECT `id`,`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`birthday`,`company_id`,`manager_id`,`active` FROM `users` WHERE age > 10 AND age < 1000 AND `users`.`deleted_at` IS NULL", tx.Statement.SQL.String())
}
