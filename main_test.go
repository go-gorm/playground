package main

import (
	"sync"
	"testing"

	"gorm.io/gorm"
)

var wg sync.WaitGroup

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

var sentences = []string{
	"Can",
	"You",
	"Please",
	"Tell",
	"What",
	"Am",
	"Doing",
	"Wrong",
}

func TestInsertInSqliteAsync(t *testing.T) {
	conn := SetupConnection()
	db, _ := conn.DB()
	defer db.Close()
	errChannel := make(chan error)
	for _, s := range sentences {
		// Implemented waitgroups because i want to log all errors
		// the channel will close after all the goroutines have run.
		wg.Add(1)
		go insert(conn, errChannel, Sentence{ActualSentence: s})
	}
	go func() {
		wg.Wait()
		close(errChannel)
	}()
	for e := range errChannel {
		t.Errorf("Error inserting in DB %+v", e)
	}
}

func insert(connection *gorm.DB, errChannel chan error, sentence Sentence) {
	defer wg.Done()
	InsertSentence(connection, errChannel, sentence)
}
