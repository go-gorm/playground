package main

import (
	"fmt"
	"testing"
)

func TestGenericsJoins(t *testing.T) {

	DB.Migrator().CreateTable(User{})

	user := User{Name: "Bob"}

	DB.Create(&user)

	var result User
	if err := DB.Table("users u").Where("u.\"name\" = ?", "Bob").First(&result); err != nil {
		fmt.Printf("failed to find user, got error: %v", err)
	}
	fmt.Printf("u.Name: %s u.ID: %d", result.Name, result.ID)
}
