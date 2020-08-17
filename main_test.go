package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

type TestTable struct {
	Address string `gorm:"unique_index:address_type_user"`
	Type    string `gorm:"unique_index:address_type_user"`
}

func TestGORM(t *testing.T) {
	DB.AutoMigrate(&TestTable{})

	var count int64
	db.Raw("select count(*) as count from pg_indexes where tablename='test_tables'").First(&count)

	if count == 0 {
		t.Errorf("create index fail")
	}
}
