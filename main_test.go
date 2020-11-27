package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

// TestPreloadGoroutine goroutine中在没有缓存的情况下使用Preload
func TestPreloadGoroutine(t *testing.T) {
	// 这里不要用Create，因为会缓存relations
	// 把db.go中init()函数中的RunMigrations()注释掉，数据库中保留一些数据来触发Preload
	var wg sync.WaitGroup

	DB = DB.Where("id = ?", 1)
	tx := DB.Session(&gorm.Session{})

	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			var user2 []User

			tx = tx.Preload("Team")
			err := tx.Find(&user2).Error

			ast := assert.New(t)
			// 这里会报错,Team: unsupported relations
			// relations是空的
			ast.Equal(nil, err)
		}()
	}
	wg.Wait()
}
