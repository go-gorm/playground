package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

const sleepMillis = 500
const testIterations = 50
const concurrency = 32

var userID uint

var errPreload = errors.New("missing related record")

func TestGORM(t *testing.T) {
	manager := User{Name: "test"}
	DB.Create(&manager)

	user := User{Name: "jinzhu", Manager: &manager}
	DB.Create(&user)
	userID = user.ID

	var wg sync.WaitGroup
	stop := make(chan struct{})
	input := make(chan context.Context)
	output := make(chan error)

	// We start some workers that will run the queries for us concurrently.
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go worker(&wg, stop, input, output, testQuery)
	}

	var failed bool

	// In each test iteration, we:
	//  - create a new context
	//  - send it to the workers, so they can start running the queries
	//  - sleep for a bit
	//  - cancel the context
	//  - collect the errors from the workers
	//
	// Each worker will stop running the queries when it encounters an error or when the context is cancelled. They will
	// then wait for another iteration, until we stop them.
	for i := 0; i < testIterations; i++ {
		ctx, cancel := context.WithCancel(context.Background())

		for j := 0; j < concurrency; j++ {
			input <- ctx
		}

		time.Sleep(sleepMillis * time.Millisecond)

		cancel()

		for j := 0; j < concurrency; j++ {
			if err := <-output; err != nil {
				t.Error(err)
				failed = true
			}
		}

		if failed {
			break
		}
	}

	close(stop)
	wg.Wait()
}

func worker(wg *sync.WaitGroup, stop <-chan struct{}, input <-chan context.Context, output chan<- error, queryFunc func(context.Context) error) {
	defer wg.Done()

	for {
		select {
		case <-stop:
			return
		case ctx := <-input:
		TEST_ITERATION:
			for {
				select {
				case <-ctx.Done():
					output <- nil
					break TEST_ITERATION
				default:
					if err := queryFunc(ctx); err != nil {
						if !errors.Is(err, context.Canceled) &&
							!strings.HasSuffix(err.Error(), " canceled") &&
							!strings.HasSuffix(err.Error(), " connection refused") &&
							!strings.HasSuffix(err.Error(), " invalid connection") {
							if ctx.Err() == context.Canceled {
								output <- err
							} else {
								output <- fmt.Errorf("the context is not cancelled but there is an error: %w", err)
							}
							break TEST_ITERATION
						}
					}
				}
			}
		}
	}
}

func testQuery(ctx context.Context) error {
	var user User
	err := DB.WithContext(ctx).Preload("Manager").Where("id = ?", userID).First(&user).Error
	if err != nil {
		return err
	}

	if user.Manager == nil {
		return errPreload
	}

	return nil
}
