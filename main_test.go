package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	for a := 0; a < 2; a++ {
		t.Run("drop and migrate", func(t *testing.T) {
			t.Run("drop", func(t *testing.T) {
				if err := DB.Migrator().DropTable(&TestTable{}); err != nil {
					t.Fatalf("failed to drop table: %v", err)
				}
			})

			for i := 0; i < 2; i++ {
				t.Run("migrate", func(t *testing.T) {
					t.Parallel()

					if err := DB.AutoMigrate(&TestTable{}); err != nil {
						t.Fatalf("failed to migrate: %v", err)
					}
				})
			}
		})
	}
}
