package main

import (
	"gorm.io/gorm"
	"sync"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

var wg sync.WaitGroup

func TestGORM(t *testing.T) {

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Point ID to 1
			model := gorm.Model{
				ID: 1,
			}

			user := User{
				Model: model,
			}
			if err := DB.FirstOrCreate(&user).Error; err != nil {
				t.Errorf("Failed, got error: %v", err)
			}
		}()
	}

	wg.Wait()
}
