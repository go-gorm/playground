package main

import (
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func WhereExpressionNotNil(t *testing.T) func(d *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		session := d.Session(&gorm.Session{})
		// show 1st clause
		if session.Statement.Clauses["WHERE"].Expression == nil {
			t.Errorf("Should not be nil")
		}
		return d
	}
}

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var userNil User
	query := DB.Scopes(WhereExpressionNotNil(t)).Where("Name NOT NULL")
	if query.Statement.Clauses["WHERE"].Expression == nil {
		t.Errorf("Should not be nil")
	}

	query.Find(&userNil)

}
