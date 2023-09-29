package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	if err := DB.Migrator().DropTable(&TestTable{}); err != nil {
		t.Fatalf("failed to drop table: %v", err)
	}
	for i := 0; i < 10; i++ {
		t.Run("test", func(t *testing.T) {
			t.Parallel()

			if err := DB.AutoMigrate(&TestTable{}); err != nil {
				t.Fatalf("failed to migrate: %v", err)
			}
		})
	}
}
