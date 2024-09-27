package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type UserOriginal struct {
	ID   uint
	Name string `gorm:"check:name_checker,name <> ''"`
}

func (UserOriginal) TableName() string {
	return "users"
}

type UserModified struct {
	ID   uint
	Name string `gorm:"check:name_checker,length(name) > 3'"`
}

func (UserModified) TableName() string {
	return "users"
}

func TestGORM(t *testing.T) {
	// Initial migration
	if err := DB.AutoMigrate(&UserOriginal{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	// Second migration
	if err := DB.AutoMigrate(&UserModified{}); err != nil {
		t.Fatalf("failed to migrate after modification: %v", err)
	}

	// Test inserting data that violates the new constraint
	user := UserModified{Name: "Hi"} // Should violate the new constraint
	err := DB.Create(&user).Error
	if err != nil {
		fmt.Printf("Expected error when inserting invalid data: %v\n", err)
	} else {
		t.Error("Unexpectedly inserted invalid data that should violate the new constraint")
	}

	// Test inserting valid data
	user = UserModified{Name: "Hello"}
	err = DB.Create(&user).Error
	if err != nil {
		t.Errorf("Error when inserting valid data: %v\n", err)
	} else {
		fmt.Println("Successfully inserted valid data")
	}

}
