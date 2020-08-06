package main

import (
	"regexp"
	"testing"

	"gorm.io/gorm"
	. "gorm.io/gorm/utils/tests"
)

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
		panic(err)	
	}
	AssertEqual(t, u.Name, "mrparano1d")
}
