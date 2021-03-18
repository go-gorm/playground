package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func LimitScope(db *gorm.DB) *gorm.DB {
	return db.Limit(1)
}

func UserScope(db *gorm.DB) *gorm.DB {
	return db.Scopes(LimitScope).Preload("Toys")
}

func TestGORM(t *testing.T) {
	users := []User{{Name: "jinzhu"}, {Name: "logan"}}

	DB.Create(&users)

	var result []User
	if err := DB.Scopes(UserScope).Find(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
