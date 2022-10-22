package main

import (
	"log"
	"strconv"
	"sync"
	"testing"
	"time"

	"gorm.io/gorm/clause"
)

const (
	count = 30
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	languages := []Language{}
	for i := 0; i < count; i++ {
		language := Language{Code: strconv.Itoa(i), Name: strconv.Itoa(i)}
		DB.Create(&language)
		languages = append(languages, language)
	}

	////////////////////////////////////////////////////////////////
	// When this is deleted, the test passes also when preloading.
	whenThisIsDeletedTheTestPasses := User{}
	DB.Create(&whenThisIsDeletedTheTestPasses)
	////////////////////////////////////////////////////////////////

	user1 := User{}
	if true { // If not preloading, the test passes.
		DB.Preload(clause.Associations).FirstOrCreate(&user1)
	} else {
		DB.FirstOrCreate(&user1)
	}

	var wg sync.WaitGroup
	for _, language := range languages {
		wg.Add(1)
		go func(userCopy User, languageCopy Language) {
			defer wg.Done()
			err := DB.Model(&userCopy).Association("Languages").Append(&languageCopy)
			if err != nil {
				log.Printf("[!] Error white appending post to user, %s", err)
				return
			}
			time.Sleep(time.Second) // Simulating task.
		}(user1, language)
	}
	wg.Wait()

	user2 := User{}
	DB.Preload(clause.Associations).First(&user2)
	log.Println(len(user2.Languages) == count, len(user2.Languages))
	if len(user2.Languages) != count {
		t.Errorf("Failed, %d should be %d", len(user2.Languages), count)
	}
}
