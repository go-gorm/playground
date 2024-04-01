package main

import (
	"context"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu", Roles: []Role{{Name: "a"}}}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	ctx := context.Background()

	var roles []Role

	err := DB.WithContext(ctx).Model(&user).Association("Roles").Find(&roles)
	if err != nil {
		panic(err)
	}
}
