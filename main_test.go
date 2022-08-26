package main

import (
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	// Check that AutoMigrate will resolve column type mismatches
	checkGoProp := "ManagerID"

	t.Logf("Type of migrator %T", DB.Migrator())

	// Get the expected type based on the dialect
	var expectedType string

	// Load the schema
	DB.Statement.Parse(&User{})
	field := DB.Statement.Schema.LookUpField(checkGoProp)

	// Get column types from particular drivers
	switch m := DB.Migrator().(type) {
	case postgres.Migrator:
		expectedType = m.DataTypeOf(field)
	case sqlite.Migrator:
		expectedType = m.DataTypeOf(field)
	case sqlserver.Migrator:
		expectedType = m.DataTypeOf(field)
	default:
		// Don't check other types
		return
	}

	expectedType = DB.Migrator().FullDataTypeOf(field).SQL
	columnName := field.DBName

	// Get type from database
	columnTypes, err := DB.Migrator().ColumnTypes(&User{})
	if err != nil {
		t.Fatalf("Failed to read column types %s", err)
	}
	var actualType string
	for _, columnType := range columnTypes {
		if columnType.Name() == columnName {
			actualType = columnType.DatabaseTypeName()
		}
	}

	// Same type?
	if expectedType != actualType {
		t.Logf("Type mismatch %s != %s, running auto migration", expectedType, actualType)
	}

	DB.AutoMigrate(&User{})

	// Reread and compare types
	columnTypes, err = DB.Migrator().ColumnTypes(&User{})
	if err != nil {
		t.Fatalf("Failed to read column types %s", err)
	}
	for _, columnType := range columnTypes {
		if columnType.Name() == columnName {
			actualType = columnType.DatabaseTypeName()
		}
	}

	if expectedType != actualType {
		t.Fatalf("AutoMigrate failed to resolve type mismatch: %s != %s. ALTER COLUMN will be called every AutoMigrate but never resolve the mismatch.", expectedType, actualType)
	}
}
