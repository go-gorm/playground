package main

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	type AutoMigrateInTx struct{}
	_ = DB.Migrator().DropTable(&AutoMigrateInTx{})
	err := DB.Transaction(func(tx *gorm.DB) error {
		err := func() error {
			type AutoMigrateInTx struct {
				First time.Time
			}
			return tx.Migrator().AutoMigrate(&AutoMigrateInTx{})
		}()
		if err != nil {
			return err
		}
		err = func() error {
			type AutoMigrateInTx struct {
				First  time.Time
				Second time.Time
			}
			return tx.Migrator().AutoMigrate(&AutoMigrateInTx{})
		}()
		if err != nil {
			return err
		}
		err = func() error {
			type AutoMigrateInTx struct {
				First  time.Time
				Second time.Time
				Third  time.Time
			}
			return tx.Migrator().AutoMigrate(&AutoMigrateInTx{})
		}()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		t.Fatalf("not expecting error, got: %v", err)
	}
	if !DB.Table("auto_migrate_in_tx").Migrator().HasColumn(&AutoMigrateInTx{}, "third") {
		t.Fatalf("expecting the third column to be present")
	}
}
