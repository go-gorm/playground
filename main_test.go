package main

import (
	"testing"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user1 := User{Name: "jinzhu", Age: 35}
	user2 := User{Name: "ortolo", Age: 36}
	user3 := User{Name: "foobar", Age: 42}

	DB.Create(&user1)
	DB.Create(&user2)
	DB.Create(&user3)

	var result User
	if err := DB.First(&result, user1.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}


	subQueryScope := func(db *gorm.DB) *gorm.DB {
		return db.Table("(?) AS sub", db.Model(&User{}))
	}

	var users []*User
	if err := DB.Scopes(subQueryScope).Find(&users).Error; err != nil {
		t.Errorf("Query with a subquery-making scope failed, got error: %v", err)
	}


	weirdSubQueryScope := func(db *gorm.DB) *gorm.DB {
		// Note the use of DB.Model(), where DB is *not* the argument
		// given to this scope but the global variable!
		return db.Table("(?) AS sub", DB.Model(&User{}))
	}

	if err := DB.Scopes(weirdSubQueryScope).Find(&users).Error; err != nil {
		t.Errorf("Query with a weird subquery-making scope failed, got error: %v", err)
	}
}
