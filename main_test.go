package main

import (
	"testing"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver


func OrderByScope(db *gorm.DB) *gorm.DB {
	return db.Order("name desc")
}

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var count int64

	t.Run("ORDER BY + Count() works fine normally", func(t *testing.T) {
		if err := DB.Model(&User{}).Order("name desc").Count(&count).Error; err != nil {
			t.Errorf("Failed in normal .Order(), got error: %v", err)
		}
		t.Logf("Count: %d", count)
	})

	t.Run("ORDER BY in scope + Count() fails", func(t *testing.T) {
		if err := DB.Model(&User{}).Scopes(OrderByScope).Count(&count).Error; err != nil {
			t.Errorf("Failed with .Order() in a scope, got error: %v", err)
		}
		t.Logf("Count: %d", count)
	})
}
