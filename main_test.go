package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	type BasePost struct {
		Id    int64
		Title string
		URL   string
	}

	type Address string

	type Author struct {
		ID      string
		Name    string
		Email   string
		Age     int
		Address Address
	}

	type HNPost struct {
		*BasePost
		Upvotes int32
		*Author `gorm:"EmbeddedPrefix:user_"` // Embedded struct
	}

	DB.Migrator().DropTable(&HNPost{})
	if err := DB.Migrator().AutoMigrate(&HNPost{}); err != nil {
		t.Fatalf("failed to auto migrate, got error: %v", err)
	}

	DB.Create(&HNPost{BasePost: &BasePost{Title: "embedded_pointer_type"}})

	var hnPost HNPost
	if err := DB.First(&hnPost, "title = ?", "embedded_pointer_type").Error; err != nil {
		t.Errorf("No error should happen when find embedded pointer type, but got %v", err)
	}

	if hnPost.Title != "embedded_pointer_type" {
		t.Errorf("Should find correct value for embedded pointer type")
	}

	if hnPost.Author != nil {
		t.Errorf("Expected to get back a nil Author but got: %v", hnPost.Author)
	}
}
