package main

import (
	"testing"

	"gorm.io/gorm"
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

func TestScopes(t *testing.T) {
	user := User{}

	sqlExpected := DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Where("name = ?", "jinzhu").Where(tx.Where("age = ?", 2).Or("age = ?", 3)).Find(&user)
	})

	scopesFunction := func(db *gorm.DB) *gorm.DB {
		return db.Where("name = ?", "jinzhu").Where(db.Where("age = ?", 2).Or("age = ?", 3))
	}

	sqlActual := DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Scopes(scopesFunction).Find(&user)
	})

	if sqlExpected != sqlActual {
		t.Errorf("Scopes() statement different. Expecting:\n%s\nbut got\n%s\n", sqlExpected, sqlActual)
	}

}

