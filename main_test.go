package main

import (
	"context"
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

	var maxId int64
	userTable := func(db *gorm.DB) *gorm.DB {
		return db.WithContext(context.Background()).Table("users")
	}
	if err := DB.Scopes(userTable).Select("max(id)").Scan(&maxId).Error; err != nil {
		t.Errorf("select max(id)")
	}
}
