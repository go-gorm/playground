package main

import (
	"strings"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Id: 1, Nickname: "user01", Avatar: "1.png", Password: "123456"}
	group := Group{Id: 1, Name: "group01", Avatar: "1.png", Password: "123456"}

	DB.Create(&user)
	DB.Create(&group)

	userInfo := &UserInfo{}
	groupInfo := &GroupInfo{}

	if err := DB.Model(&User{}).First(userInfo).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Model(&Group{}).First(groupInfo).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if !strings.HasPrefix(groupInfo.Avatar, ossUrl) {
		t.Error("Failed")
	}

	if !strings.HasPrefix(userInfo.Avatar, ossUrl) {
		t.Error("Failed")
	}
}
