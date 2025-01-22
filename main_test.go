package main

import (
	"context"
	"sync"
	"testing"
	"time"
)

const concurrentReads = 40

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	testRunSuccessful := false
	wgSuccess := sync.WaitGroup{}
	wgSuccess.Add(concurrentReads)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	start := make(chan struct{})

	for i := 0; i < concurrentReads/2; i++ {
		go func() {
			t.Logf("Entered routine 1-%d", i)
			var result User

			<-start
			transaction := DB.Begin()
			if err := transaction.First(&result, "id = ? ", 1).Error; err != nil {
				transaction.Rollback()
				return
			}
			transaction.Commit()
			t.Log("Got User from routine 1")
			wgSuccess.Done()
		}()
	}

	for i := 0; i < concurrentReads/2; i++ {
		go func() {
			t.Logf("Entered routine 2-%d", i)
			var result User

			<-start
			if err := DB.First(&result, "id = ? ", 1).Error; err != nil {
				t.Errorf("Failed, got error: %v", err)
				return
			}
			t.Log("Got User from routine 2")
			wgSuccess.Done()
		}()
	}

	time.Sleep(200 * time.Millisecond)
	close(start)
	t.Log("Started routines")

	go func() {
		wgSuccess.Wait()
		testRunSuccessful = true
	}()

	<-ctx.Done()
	if !testRunSuccessful {
		t.Fatalf("Test failed")
	}

	t.Logf("Test completed")
}
