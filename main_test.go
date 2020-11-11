package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type TestUniqueIdx struct {
	gorm.Model
	Foo string `form:"foo" json:"foo" gorm:"unique_index:idx_foo_bar"`
	Bar string `form:"bar" json:"bar" gorm:"unique_index:idx_foo_bar"`
	Baz int    `form:"baz" json:"baz"  gorm:"default:500"`
}

func addTestUniqueIdx(test TestUniqueIdx) error {
	return DB.Create(&test).Error
}

func updateTestUniqueIdxMap(test map[string]interface{}) error {
	return DB.Model(TestUniqueIdx{}).Where("foo = ?", test["foo"]).Updates(test).Error
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
		Baz: 100,
	}
	testMap := make(map[string]interface{})
	testMap["foo"] = foo
	testMap["bar"] = bar
	testMap["baz"] = 100

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
