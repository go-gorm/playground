package main

import (
	"testing"

	"github.com/google/uuid"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Post struct {
	ID         uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Title      string
	Categories []*Category `gorm:"Many2Many:posts_categories"`
}

type Category struct {
	ID    uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Title string
	Posts []*Post `gorm:"Many2Many:posts_categories"`
}

func TestGORM(t *testing.T) {
	post := Post{
		Title: "Hello World",
		Categories: []*Category{
			{Title: "Coding"},
			{Title: "Golang"},
		},
	}

	if err := DB.Create(&post).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
