package main

import (
	"sync"
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

// TestBug_DecimalSizeNilDerefOnRace reproduces a nil pointer dereference at
// gorm.io/gorm@v1.31.1/migrator/column_type.go:75 (ColumnType.DecimalSize),
// reached through migrator/migrator.go:528 inside (Migrator).MigrateColumn.
//
// Root cause: the postgres driver builds ColumnType wrappers in two passes —
//   - pass 1: information_schema.columns
//   - pass 2: SELECT * FROM <table> LIMIT 1, then assign each *sql.ColumnType
//     into the matching wrapper's named SQLColumnType field.
//
// A concurrent DROP COLUMN landing between the two passes leaves a wrapper
// for the dropped column with SQLColumnType == nil. The next call to a
// delegating wrapper method (DecimalSize/Length/Nullable/ScanType) panics.
//
// A nil-check at the call site does NOT help — the wrapper *value* is non-nil;
// the nil pointer is the named SQLColumnType field inside it. The fix must
// live in the wrapper methods.
//
// Source paths:
//   - https://github.com/go-gorm/gorm/blob/v1.31.1/migrator/column_type.go#L8-L24
//   - https://github.com/go-gorm/gorm/blob/v1.31.1/migrator/column_type.go#L71-L76
//   - https://github.com/go-gorm/gorm/blob/v1.31.1/migrator/migrator.go#L528
//   - https://github.com/go-gorm/postgres/blob/v1.6.0/migrator.go#L522-L540
//
// Postgres-specific (depends on the driver's two-pass ColumnTypes shape).
func TestBug_DecimalSizeNilDerefOnRace(t *testing.T) {
	if DB.Dialector.Name() != "postgres" {
		t.Skip("driver-shape-specific bug; skipping on", DB.Dialector.Name())
	}

	type Supplier struct {
		ID       uint   `gorm:"primaryKey"`
		Category string `gorm:"type:varchar(50);not null"`
	}

	t.Cleanup(func() { DB.Exec(`DROP TABLE IF EXISTS suppliers`) })

	const iterations = 50
	var (
		panics int
		mu     sync.Mutex
	)

	for attempt := 1; attempt <= iterations; attempt++ {
		// Reset to a known starting state: table with a "category" column
		// of width 20. Model wants width 50 → AutoMigrate will need to alter.
		DB.Exec(`DROP TABLE IF EXISTS suppliers`)
		DB.Exec(`CREATE TABLE suppliers (id bigserial PRIMARY KEY, category varchar(20) NOT NULL)`)

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					mu.Lock()
					panics++
					mu.Unlock()
				}
			}()
			_ = DB.AutoMigrate(&Supplier{})
		}()

		go func() {
			defer wg.Done()
			defer func() { _ = recover() }()
			DB.Exec(`ALTER TABLE suppliers DROP COLUMN IF EXISTS category`)
		}()

		wg.Wait()
	}

	t.Logf("ran %d race iterations, captured %d nil-deref panic(s)", iterations, panics)
	if panics == 0 {
		t.Fatal("expected at least one nil-deref panic from the AutoMigrate goroutine, got none — bug appears patched")
	}
}
