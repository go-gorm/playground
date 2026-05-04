package main

import (
	"testing"

	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

func TestGORM(t *testing.T) {
	user := models.User{Name: "jinzhu"}

	DB.Create(&user)

	var result models.User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

// TestBug_AlterColumnNilDerefOnMissingColumn reproduces a nil pointer
// dereference at gorm.io/driver/postgres@v1.6.0/migrator.go:336.
//
// Trigger condition (deterministic, single call):
//   - The schema declares a field whose DBName is absent from the DB table.
//   - (Migrator).AlterColumn() loops over ColumnTypes() result, never assigns
//     fieldColumnType, then dereferences it on the next line — panic.
//
// Source path: https://github.com/go-gorm/postgres/blob/v1.6.0/migrator.go#L319-L336
//
// This test is postgres-specific: skip on other dialects.
func TestBug_AlterColumnNilDerefOnMissingColumn(t *testing.T) {
	if DB.Dialector.Name() != "postgres" {
		t.Skip("postgres-driver-specific bug; skipping on", DB.Dialector.Name())
	}

	type Supplier struct {
		ID       uint   `gorm:"primaryKey"`
		Category string `gorm:"type:varchar(20);not null"`
	}

	// Seed a table with ONLY "id" — the schema field "Category" (DBName
	// "category") will not match any DB column.
	if err := DB.Exec(`DROP TABLE IF EXISTS suppliers`).Error; err != nil {
		t.Fatalf("drop: %v", err)
	}
	if err := DB.Exec(`CREATE TABLE suppliers (id bigserial PRIMARY KEY)`).Error; err != nil {
		t.Fatalf("create: %v", err)
	}
	t.Cleanup(func() { DB.Exec(`DROP TABLE IF EXISTS suppliers`) })

	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected nil-deref panic at driver/postgres migrator.go:336, got none — bug appears patched")
		}
		t.Logf("reproduced expected panic: %v", r)
	}()

	// LookUpField("Category") returns non-nil (field exists in schema).
	// ColumnTypes() returns only "id" (column "category" doesn't exist).
	// Loop never assigns fieldColumnType. Next deref panics at line 336.
	_ = DB.Migrator().AlterColumn(&Supplier{}, "Category")
}
