package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

// Both models declare a unique index with the same name. PostgreSQL (and
// SQLite) keep index names unique per schema, so the two indexes cannot both
// exist under that name. AutoMigrate should either create both or return an
// error; on PostgreSQL it returns nil and silently skips the second one.
type SharedIndexLeft struct {
	ID   uint
	Host string `gorm:"size:64;uniqueIndex:idx_shared_host_port"`
	Port int    `gorm:"uniqueIndex:idx_shared_host_port"`
}

type SharedIndexRight struct {
	ID   uint
	Host string `gorm:"size:64;uniqueIndex:idx_shared_host_port"`
	Port int    `gorm:"uniqueIndex:idx_shared_host_port"`
}

func TestGORM(t *testing.T) {
	if err := DB.Migrator().DropTable(&SharedIndexLeft{}, &SharedIndexRight{}); err != nil {
		t.Fatalf("drop tables: %v", err)
	}
	if err := DB.AutoMigrate(&SharedIndexLeft{}, &SharedIndexRight{}); err != nil {
		t.Fatalf("AutoMigrate returned an error: %v", err)
	}
	for _, model := range []any{&SharedIndexLeft{}, &SharedIndexRight{}} {
		if !DB.Migrator().HasIndex(model, "idx_shared_host_port") {
			t.Errorf("AutoMigrate returned nil, but the unique index declared on %T was not created", model)
		}
	}
	if err := DB.Create(&SharedIndexRight{Host: "a", Port: 1}).Error; err != nil {
		t.Fatalf("first insert: %v", err)
	}
	if err := DB.Create(&SharedIndexRight{Host: "a", Port: 1}).Error; err == nil {
		t.Errorf("a duplicate (Host, Port) row was accepted although the model declares a unique index")
	}
}
