package main

import (
	"encoding/json"
	"gorm.io/gorm"
	"testing"
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

func TestDeletedAtUnMarshal(t *testing.T) {

	expected := &gorm.Model{}
	b, _ := json.Marshal(expected)

	result := &gorm.Model{}
	_ = json.Unmarshal(b, result)
	if result.DeletedAt != expected.DeletedAt {
		t.Errorf("Failed, result.DeletedAt: %v is not same as expected.DeletedAt: %v", result.DeletedAt, expected.DeletedAt)
	}
}
