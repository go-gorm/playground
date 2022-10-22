package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	for i := 0; i < 10; i++ {
		log.Println(i)
		language := Language{Code: fmt.Sprintf("Code%d", i), Name: fmt.Sprintf("Name%d", i)}

		start := make(chan struct{})
		var wg sync.WaitGroup
		for j := 0; j < 20; j++ {
			wg.Add(1)
			go func(languageCopy Language) {
				defer wg.Done()
				<-start
				tx := DB.FirstOrCreate(&languageCopy)
				if tx.Error != nil && strings.Contains(tx.Error.Error(), "UNIQUE constraint failed") {
					t.Error(tx.Error.Error())
					os.Exit(1)
				}
			}(language)
		}
		close(start)
		wg.Wait()
	}
}
