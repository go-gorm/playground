package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	rdb := DB.Model(&User{}).Select("name").First(&User{})
	if rdb.Error != nil {
		t.Errorf("select `name` should success")
		t.Fail()
	}

	rdb = DB.Model(&User{}).Select("name, age").First(&User{})
	if rdb.Error != nil {
		t.Errorf("select `name, age` should success")
		t.Fail()
	}

	rdb = DB.Model(&User{}).Select("name as n, age as a").First(&User{})
	if rdb.Error != nil {
		t.Errorf("select `name as n, age as a` should success")
		t.Fail()
	}

	rdb = DB.Model(&User{}).Select("name as n").First(&User{})
	if rdb.Error == nil {
		t.Errorf("select `name as n` should return error")
		t.Fail()
	}
}
