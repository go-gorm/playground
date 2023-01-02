package main

import (
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := SoftDeleteUser{Name: "jinzhu"}
	DB.AutoMigrate(SoftDeleteUser{})
	DB.Create(&user)

	var results []SoftDeleteUser
	// No error when table alias isn't used.
	if err := DB.Table("soft_delete_users").Where("id = ?", 1).Find(&results).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// Error occurs when table alias is used due to soft delete behavior using the name of the table as the alias on
	// the soft delete column.
	if err := DB.Table("soft_delete_users sds").Where("sds.id = ?", 1).Find(&results).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

type SoftDeleteUser struct {
	gorm.Model
	Name    string
	Deleted gorm.DeletedAt
}
