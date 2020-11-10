package main

import (
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Model struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

type TestUniqueIdx struct {
	// Model
	Foo string `form:"foo" json:"foo" gorm:"unique_index:idx_foo_bar"`
	Bar string `form:"bar" json:"bar" gorm:"unique_index:idx_foo_bar"`
}

func addTestUniqueIdx(test TestUniqueIdx) error {
	return DB.Create(&test).Error
}

func updateTestUniqueIdxMap(test map[string]interface{}) error {
	return DB.Model(&TestUniqueIdx{}).Where("foo = ?", test["foo"]).Updates(test).Error
}

func updateTestUniqueIdx(test TestUniqueIdx) error {
	return DB.Model(&TestUniqueIdx{}).Where("foo = ?", test.Foo).Updates(test).Error
}

func saveTestUniqueIdx(test TestUniqueIdx) error {
	return DB.Model(&TestUniqueIdx{}).Where("foo = ?", test.Foo).Save(&test).Error
}

func TestGORM(t *testing.T) {
	foo := "jhbcxwc"
	bar := "wbwkxccw"
	test := TestUniqueIdx{
		Foo: foo,
		Bar: bar,
	}
	testMap := make(map[string]interface{})
	testMap["foo"] = foo
	testMap["bar"] = bar

	if err := DB.AutoMigrate(&TestUniqueIdx{}); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := addTestUniqueIdx(test); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := updateTestUniqueIdxMap(testMap); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := updateTestUniqueIdx(test); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := saveTestUniqueIdx(test); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
