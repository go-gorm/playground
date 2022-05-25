package main

import (
	"testing"

)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	// Migrate the schema
	_ = DB.AutoMigrate(&User{})
	_ = DB.AutoMigrate(&UserPermission{})

	//l1 := &UserPermission{UserId: "Chinese", Name: "可能是中文"}
	//DB.Create(l1)

	var up UserPermission
	err := DB.Where(&UserPermission{UserId: "jinzhu",} ).Find(&up).Error // find language with code Chinese
	if err != nil {
		t.Errorf("err is %v, expect : nil", err)
	}
}
