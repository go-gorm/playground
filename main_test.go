package main

import (
	"context"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// test
	portal := PortalDb{
		Name: "简单测试",
	}

	DB.Create(&portal)

	if err := portal.Get(context.Background()); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
