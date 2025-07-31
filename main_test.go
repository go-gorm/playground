package main

import (
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
}

func TestTableNameWithDash(t *testing.T) {
	table := &TableNameWithDash{}

	if err := DB.Migrator().DropTable(table); err != nil {
		t.Fatalf("Failed to drop table, got error %v\n", err)
	}

	if err := DB.AutoMigrate(&TableNameWithDash{}); err != nil {
		t.Fatalf("Failed to migrate, got error %v\n", err)
	}
	if err := DB.AutoMigrate(&TableNameWithDash{}); err != nil {
		t.Errorf("Failed to re-migrate, got error %v\n", err)
	}
}
