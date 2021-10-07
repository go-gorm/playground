package main

import (
	"fmt"
	"strings"
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type NewUser struct {
	gorm.Model
	Name  string `gorm:"index"`
	City  string `gorm:"index"`
	Table string `gorm:"-"`
}

/*
NOTE: https://gorm.io/docs/v2_release_note.html#Breaking-Changes
*/
func UserTable(u *NewUser) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table(fmt.Sprintf("newuser_%s", strings.ToLower(u.Table)))
	}
}

// TableName which is cached
func (u *NewUser) TableName() string {
	return fmt.Sprintf("newuser_%s", u.Table)
}

func TestGORM(t *testing.T) {
	tableName := []string{"a", "b", "c"}
	for _, v := range tableName {
		nu := &NewUser{
			Table: v,
		}
		DB.Scopes(UserTable(nu)).Migrator().AutoMigrate(&nu)
		if err := DB.Scopes(UserTable(nu)).Migrator().AutoMigrate(&nu); err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
	}
}
