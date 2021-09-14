package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sentence struct {
	gorm.Model
	ActualSentence string
}

func main() {
	fmt.Print("Hello\n")
}

func SetupConnection() *gorm.DB {
	connection, connectionError := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if connectionError != nil {
		log.Panic("Cannot create a database connection", connectionError)
	}
	connection.AutoMigrate(&Sentence{})
	return connection
}

func InsertSentence(connection *gorm.DB, errChannel chan error, sentence Sentence) {
	res := connection.Create(&sentence)
	if res.Error != nil {
		errChannel <- res.Error
	} else {
		fmt.Printf("Successfully inserted %s\n", sentence.ActualSentence)
	}
}
