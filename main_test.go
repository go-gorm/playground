package main

import (
	"testing"
	"time"

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

type TimestampTest struct {
	TimeAt *time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
}

func TestMigrateTimestamp(t *testing.T) {
	err := DB.AutoMigrate(TimestampTest{})
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

	migrator.RunWithValue(&TimestampTest{}, func(stmt *gorm.Statement) (err error) {
		field, ok := stmt.Schema.FieldsByDBName["time_at"]
		if !ok {
			t.Errorf("time_at not found")
		} else if field.DefaultValueInterface == nil {
			t.Errorf("expected field.DefaultValueInterface to be non-nil: %+v", field)
		}

		return nil
	})
}
