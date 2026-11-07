package main

import (
	"testing"

	"gorm.io/gorm"
	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

// PanicOnCreate panics in its BeforeCreate hook, as any hook with a bug can.
type PanicOnCreate struct {
	ID   uint
	Name string
}

func (*PanicOnCreate) BeforeCreate(*gorm.DB) error {
	panic("BeforeCreate failed")
}

// A panic inside Create, from GORM itself (a nil element in an association
// slice) or from a hook, leaves the default transaction open: the callback that
// commits or rolls it back never runs, so the connection is never returned to
// the pool. DB.Transaction rolls back when its function panics; the default
// transaction should do the same before the panic propagates.
func TestGORM(t *testing.T) {
	if err := DB.AutoMigrate(&PanicOnCreate{}); err != nil {
		t.Fatalf("AutoMigrate: %v", err)
	}
	sqlDB, err := DB.DB()
	if err != nil {
		t.Fatalf("DB.DB: %v", err)
	}
	for name, create := range map[string]func(){
		"nil element in a has-many slice": func() { DB.Create(&models.User{Name: "nil-pet", Pets: []*models.Pet{nil}}) },
		"panicking BeforeCreate hook":     func() { DB.Create(&PanicOnCreate{Name: "hook"}) },
	} {
		before := sqlDB.Stats().InUse
		recovered := func() (value any) {
			defer func() { value = recover() }()
			create()
			return nil
		}()
		if inUse := sqlDB.Stats().InUse; inUse != before {
			t.Errorf("%s: Create panicked (%v) and left the default transaction open: %d connection(s) in use, want %d", name, recovered, inUse, before)
		}
	}
}
