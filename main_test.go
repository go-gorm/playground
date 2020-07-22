package main

import (
	"testing"
)

type _User1 struct {
	Id        uint64         `gorm:"primarykey"`
	Name      string         `gorm:"type:varchar(255);uniqueIndex:uk_name_delete_at"`
	DeletedAt gorm.DeletedAt `gorm:"uniqueIndex:uk_name_delete_at"`
	gorm.Model
}

type _User2 struct {
	Id        uint64         `gorm:"primarykey"`
	Name      string         `gorm:"type:varchar(255);uniqueIndex:uk_name_delete_at"`
	DeletedAt gorm.DeletedAt `gorm:"uniqueIndex:uk_name_delete_at;default:'1999-01-01'"`
	gorm.Model
}

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func Test_XXX(t *testing.T) {
	DB.AutoMigrate(&_User1{})
	DB.AutoMigrate(&_User2{})

	// without delete_at
	DB.Create(&_User1{Name: "user1"})
	rdb := DB.Create(&_User1{Name: "user1"})
	if rdb.Error == nil {
		// for null not support unique key
		t.Error("DB.Create(user) should be duplicate entry for key 'uk_name_delete_at'")
	}

	// with delete_at
	DB.Create(&_User2{Name: "user2"})
	DB.Model(&_User2{}).Where(&_User2{Name: "user2"}).Delete(&User{Name: "user2"})
	rdb = DB.Create(&_User2{Name: "user2"})
	if rdb.Error != nil {
		// the first user is delete, the next user use the first user's name
		t.Error("rdb.Error should be nil")
	}
	user := &_User2{}
	rdb = DB.Where(&_User2{Name: "user2"}).First(user)
	if rdb.RowsAffected == 0 {
		// For default value of delete_at field doesn't work
		t.Error("rdb.RowsAffected should be 1")
	}
}
