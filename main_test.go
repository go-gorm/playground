package main

import (
	"gorm.io/gorm"
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
	t.Logf("%+v\n",user)
	if err := DB.Where("name = ?", "jinzhu").UpdateColumn("age", gorm.Expr("age + ?", 1)).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
