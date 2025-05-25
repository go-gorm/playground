package main

import (
	"sync"
	"testing"
	"time"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	db := DB

	transactionalClients := []string{"Alice", "Bob"}
	var started, finished sync.WaitGroup
	started.Add(len(transactionalClients))
	finished.Add(len(transactionalClients))

	for _, clientName := range transactionalClients {
		go func() {
			db.Transaction(func(db *gorm.DB) error {
				started.Done()
				defer finished.Done()

				t.Logf("%s has a connection and is thinking...", clientName)
				doSomeWork()

				return addTwoNumbers(t, db, clientName)
			})
		}()
	}
	started.Wait()

	addTwoNumbers(t, db, "Camille")
	finished.Wait()
}

func addTwoNumbers(t *testing.T, db *gorm.DB, clientName string) error {
	t.Logf("%s wants to add two numbers...", clientName)
	var sum int32
	err := db.Raw("SELECT $1::int + $2::int", 2, 2).First(&sum).Error
	if err != nil {
		t.Errorf("%s failed: %s", clientName, err)
		return err
	}
	t.Logf("%s thinks that 2 + 2 = %v", clientName, sum)
	return nil
}

func doSomeWork() {
	time.Sleep(time.Second)
}
