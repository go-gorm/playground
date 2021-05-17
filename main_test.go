package main

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func TestScope(t *testing.T) {
	const name = "TestScope"

	// prepare an User
	user := User{Name: name}
	if err := DB.Create(&user).Error; err != nil {
		t.Errorf("DB create failed, err=%s", err.Error())
	}

	ctx := context.TODO() // some context
	// whereScope can be reused in a Find & a Count, a common usage in paginate APIs.
	whereScope := func(db *gorm.DB) *gorm.DB {
		return db.WithContext(ctx).Table("users").Where(clause.Eq{Column: "name", Value: name})
	}

	var total int64
	if err := DB.Scopes(whereScope).Count(&total).Error; err != nil {
		t.Errorf("DB count failed, err=%s", err.Error())
	}
	if total != 1 { // should find a record, but count always return 0
		t.Errorf("DB count failed, want=%d, got=%d", 1, total)
	}
}
