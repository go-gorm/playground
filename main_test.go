package main

import (
	"testing"

	"gorm.io/gorm/clause"
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

func TestChainingShouldNotAffectPreviouslySavedInstances(t *testing.T) {
	db1 := DB.Table("users")
	db2 := db1.Table("sessions")
	if db1 == db2 {
		t.Errorf("db1 should not be equal to db2")
	}
	db3 := db2.Where("1=1")
	if db2 == db3 {
		t.Errorf("db2 should not be equal to db3")
	}
	db4 := db3.Select("*")
	if db3 == db4 {
		t.Errorf("db3 should not be equal to db4")
	}
	db5 := db4.Joins("JOIN a USING(id)")
	if db4 == db5 {
		t.Errorf("db4 should not be equal to db5")
	}
	db6 := db5.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("amount > ?", 1000)
	})
	if db5 == db6 {
		t.Errorf("db5 should not be equal to db6")
	}
	db7 := db6.Preload("Orders", "state NOT IN (?)", "cancelled")
	if db6 == db7 {
		t.Errorf("db6 should not be equal to db7")
	}
	db8 := db7.Raw("SELECT abcd")
	if db7 == db8 {
		t.Errorf("db7 should not be equal to db8")
	}
	db9 := db8.Omit("id")
	if db8 == db9 {
		t.Errorf("db8 should not be equal to db9")
	}
	db10 := db9.Model(&User{})
	if db9 == db10 {
		t.Errorf("db9 should not be equal to db10")
	}
	db11 := db10.Clauses(clause.Expr{})
	if db10 == db11 {
		t.Errorf("db10 should not be equal to db11")
	}
	db12 := db11.Distinct()
	if db11 == db12 {
		t.Errorf("db11 should not be equal to db12")
	}
	db13 := db12.Not("1")
	if db12 == db13 {
		t.Errorf("db12 should not be equal to db13")
	}
	db14 := db13.Or("1")
	if db13 == db14 {
		t.Errorf("db13 should not be equal to db14")
	}
	db15 := db14.Group("id")
	if db14 == db15 {
		t.Errorf("db14 should not be equal to db15")
	}
	db16 := db15.Having("1")
	if db15 == db16 {
		t.Errorf("db15 should not be equal to db16")
	}
	db17 := db16.Limit(1)
	if db16 == db17 {
		t.Errorf("db16 should not be equal to db17")
	}
	db18 := db17.Offset(2)
	if db17 == db18 {
		t.Errorf("db17 should not be equal to db18")
	}
	db19 := db18.Attrs(User{Age: 20})
	if db18 == db19 {
		t.Errorf("db18 should not be equal to db19")
	}
	db20 := db19.Assign(User{Age: 20})
	if db19 == db20 {
		t.Errorf("db19 should not be equal to db20")
	}
	db21 := db20.Unscoped()
	if db20 == db21 {
		t.Errorf("db20 should not be equal to db21")
	}

	var result []map[string]interface{}
	if err := db1.Scan(&result).Error; err != nil {
		t.Errorf("error on selecting * from users: %v", err)
	}
}
