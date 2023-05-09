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

func TestMigrateExistingBoolColumnPG(t *testing.T) {
	if DB.Dialector.Name() != "postgres" {
		return
	}

	type ColumnStruct struct {
		BooleanColumn int `gorm:"type:smallint;default:0"`
	}

	type ColumnStruct2 struct {
		BooleanColumn bool `gorm:"default:false"` // change existing boolean column from smallint with default 0 to boolean with default false
	}

	DB.Migrator().DropTable(&ColumnStruct{})

	if err := DB.AutoMigrate(&ColumnStruct{}); err != nil {
		t.Errorf("Failed to migrate, got %v", err)
	}

	DB.Table("column_structs").AutoMigrate(&ColumnStruct2{}) // expect no error
}
