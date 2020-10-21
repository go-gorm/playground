package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	stop := make(chan struct{})
	wg := sync.WaitGroup{}

	n := 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go dbconsumer(t, DB, stop, &wg)
	}

	time.Sleep(2 * time.Second)

	close(stop)
	wg.Wait()
}

func dbconsumer(t *testing.T, db *gorm.DB, stop <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	var result struct {
		N int64
	}
	var err error
	query := "select 42 as n"

	for {
		select {
		case <-stop:
			return
		default:
			r := rand.Intn(100)
			time.Sleep(time.Duration(r) * time.Microsecond)

			// If we do this, there is no race.
			// db = db.WithContext(context.Background())

			err = db.Raw(query).Scan(&result).Error
			if err != nil {
				t.Error(err)
				return
			}
			if result.N != 42 {
				t.Errorf("expected result to be 42, got: %d", result.N)
				return
			}

			result.N = 0
		}
	}
}
