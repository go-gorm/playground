package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type UserWithTable struct {
	gorm.Model
	Name string
}

func (UserWithTable) TableName() string {
	return "gorm.user"
}

func init() {
	DB.AutoMigrate(&UserWithTable{})
}

func TestTable(t *testing.T) {

	var u UserWithTable
	u.Name = "mrparano1d"

	if err := DB.Create(&u).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
