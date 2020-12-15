package main

import (
	"fmt"
	"sync"
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	wg := sync.WaitGroup{}
	jobChan := make(chan bool, 200)

	for i := 0; i < 1000; i++ {
		jobChan <- true
		wg.Add(1)
		go func() {
			user := User{Name: "jinzhu"}

			db, _ := DB.DB()
			DB.Transaction(func(tx *gorm.DB) error {
				fmt.Printf("%+v\n", db.Stats())
				return tx.Create(&user).Error
			})
			wg.Done()
			<-jobChan
		}()
	}
	wg.Wait()

}
