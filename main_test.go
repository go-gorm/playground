package main

import (
	"database/sql"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}
	DB.Create(&user)

	account := Account{UserID: sql.NullInt64{Int64: int64(user.ID), Valid: true}}
	DB.Create(&account)

	DB.Delete(&user)
}
