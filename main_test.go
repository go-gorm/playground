package main

import (
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	// setup
	var count int64
	DB.Model(&User{}).Count(&count)
	if count == 0 {
		DB.Create(&User{Name: "name"})
		DB.Create(&Account{UserID: sql.NullInt64{Int64: 1}})
	}

	// main logic
	var users []User
	var accounts []Account
	var i int64

	DB = DB.Where("name = ?", "name")
	err := DB.Order("id DESC").Limit(10).Offset(0).Find(&users).Error
	if err != nil {
		require.Nil(t, err)
	}
	err = DB.Model(&User{}).Count(&i).Error
	if err != nil {
		require.Nil(t, err)
	}
	err = DB.Find(&accounts).Error
	require.Nil(t, err)
}
