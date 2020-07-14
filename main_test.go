package main

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"runtime"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func gen() *gorm.DB {
	return DB.Model(&User{}).Where(&User{Age: 100})
}

func TestGORM(t *testing.T) {
	DB.Create(&User{Name: "User 1", Age: 1})
	DB.Create(&User{Name: "User 2", Age: 10})
	DB.Create(&User{Name: "User 3", Age: 100})
	DB.Create(&User{Name: "User 4", Age: 1000})

	tx := gen()

	runtime.GC()

	tx = tx.Session(&gorm.Session{})

	var c int64
	err := tx.Count(&c).Error // will trigger `unsupported data type: <P>`
	assert.NoError(t, err)
	assert.Equal(t, int64(1), c)
}
