package main

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/migrator"
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

type UniqueTest struct {
	UniqueIndex string `gorm:"size:191;uniqueIndex"`
}

func TestMigrateUnique(t *testing.T) {
	err := DB.AutoMigrate(UniqueTest{})
	if err != nil {
		t.Errorf("AutoMigrate failed: %v", err)
	}

	var migrator migrator.Migrator
	switch v := DB.Migrator().(type) {
	case sqlite.Migrator:
		migrator = v.Migrator
	case mysql.Migrator:
		migrator = v.Migrator
	case postgres.Migrator:
		migrator = v.Migrator
	case sqlserver.Migrator:
		migrator = v.Migrator
	default:
		t.Fatalf("unrecognized type: %T", v)
	}

	migrator.RunWithValue(&UniqueTest{}, func(stmt *gorm.Statement) (err error) {
		field, ok := stmt.Schema.FieldsByDBName["unique_index"]
		if !ok {
			t.Errorf("unique_index not found")
		} else if !field.Unique {
			t.Errorf("expected field.unique to be true, got false: %+v", field)
		}

		return nil
	})
}
