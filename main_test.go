package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	post := Post{
		Title: "Hello World",
		Content:
		"Foo Bar",
		Threads: []Thread{
			{
				Title:   "Foo",
				Content: "Bar",
			},
			{
				Title:   "Foo21",
				Content: "Bar1",
			},
			{
				Title:   "Foo2",
				Content: "Bar2",
			},
		},
	}

	if err := DB.Create(&post).Error; err != nil {
		t.Errorf("Failed to create post, got error: %v", err)
	}

	if err := DB.Delete(&post).Error; err != nil {
		t.Errorf("Failed to delete post, got error: %v", err)
	}

	var threads []Thread
	if err := DB.Find(&threads).Error; err != nil {
		t.Errorf("Failed to fetch threads, got error: %v", err)
	}

	if len(threads) > 0 {
		t.Errorf("Failed to delete all threads, given parent post is deleted")
	}
}
