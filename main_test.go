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

func TestMigrateExistingSmallintBoolColumnPG(t *testing.T) {
	if DB.Dialector.Name() != "postgres" {
		return
	}

	type ColumnStruct struct {
		Name     string
		IsActive bool `gorm:"type:smallint"`
	}

	type ColumnStruct2 struct {
		Name     string
		IsActive bool // change existing boolean column from smallint or other to boolean
	}

	DB.Migrator().DropTable(&ColumnStruct{})

	if err := DB.AutoMigrate(&ColumnStruct{}); err != nil {
		t.Errorf("Failed to migrate, got %v", err)
	}

	if err := DB.Table("column_structs").AutoMigrate(&ColumnStruct2{}); err != nil {
		t.Fatalf("no error should happened when auto migrate column, but got %v", err)
	}
}
