package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Book struct {
	BookID int    `gorm:"column:book_id;autoIncrement:true;autoIncrementIncrement:10"`
	Name   string `gorm:"column:name;type:varchar(255)"`
}

func TestGORM(t *testing.T) {
	err := DB.AutoMigrate(&Book{})
	if err != nil {
		t.Error(err.Error())
	}
	var increment int
	err = DB.Raw(`select increment from information_schema.sequences where sequence_name = 'books_book_id_seq'`).
		Scan(&increment).Error
	if err != nil {
		t.Error(err.Error())
	}

	// should be 10, but actually get 1
	if increment != 10 {
		t.Error("auto increment should be 10, but get", increment)
	}
}
