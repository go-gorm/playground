package main

import (
	"testing"
	
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Message struct {
	gorm.Model
	Title	string
	Body	string
	Slug	string `gorm:"unique;not null"`
}

func TestGORM(t *testing.T) {
	message = Message{Title: "", Body: "", Slug: ""}

	err := DB.Create(&message).Error.(*pq.Error)

	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
