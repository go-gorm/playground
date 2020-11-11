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
	CreatedAt time.Time  `json:"created_at" gorm:"index"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

type Test struct {
	Model
	Foo string `form:"foo" json:"foo" gorm:"uniqueIndex:idx_tests_foo_bar"`
	Bar string `form:"bar" json:"bar" gorm:"uniqueIndex:idx_tests_foo_bar"`
	Baz int    `form:"baz" json:"baz"  gorm:"default:500"`
}

func TestGORM(t *testing.T) {
	foo := "jhbcxwc"
	bar := "wbwkxccw"
	test := Test{
		Foo: foo,
		Bar: bar,
		Baz: 100,
	}
	testMap := make(map[string]interface{})
	testMap["foo"] = foo
	testMap["bar"] = bar
	testMap["baz"] = 100

	if err := DB.AutoMigrate(&Test{}); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if hasIndex := DB.Migrator().HasIndex(&Test{}, "idx_tests_created_at"); !hasIndex {
		t.Errorf("Fail to create index idx_tests_created_at")
	}

	if hasIndex := DB.Migrator().HasIndex(&Test{}, "idx_tests_foo_bar"); !hasIndex {
		t.Errorf("Fail to create index idx_tests_foo_bar")
	}

	if err := DB.Create(&test).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Where("foo = ?", test.Foo).Delete(&test).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Model(&test).Create(testMap).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Model(Test{}).Where("foo = ?", testMap["foo"]).Updates(testMap).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Model(&Test{}).Where("foo = ?", test.Foo).Updates(test).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Model(&Test{}).Where("foo = ?", test.Foo).Save(&test).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
