package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

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

// import csv data.
func importData() {
	dbFile := filepath.Join(os.TempDir(), "gorm.db")
	os.Remove(dbFile)

	file, err := os.OpenFile(".init.sql", os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	modeCmd := fmt.Sprintln(".mode csv")
	if _, err := file.Write([]byte(modeCmd)); err != nil {
		log.Fatalf("error while writing the file. %v", err)
	}

	importCmd := fmt.Sprintf(".import %v %v \n", "data.csv", "data")
	file.WriteString(importCmd)

	// import to database.
	exec.Command("sqlite3", "--init", file.Name(), filepath.Join(os.TempDir(), "gorm.db")).Run()

	if err := os.Remove(file.Name()); err != nil {
		log.Fatalf("error while removing the file. %v", err)
	}

}

func TestGORMTime(t *testing.T) {

	importData()

	var result []Data

	// correct way query.
	// DB.Table("data").Where("start_time >= ?", "2022-07-01 00:00:00").Find(&result)

	// error way query.
	time, _ := time.Parse("2006-01-02 15:04:05", "2022-07-01 00:00:00")
	DB.Table("data").Where("start_time >= ?", time).Find(&result)

	if len(result) != 2 {
		t.Errorf("Failed, got %v", len(result))
	}

}
