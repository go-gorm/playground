package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	// Drop index failure
	// err := DB.Migrator().DropIndex(&User{}, "name_idx")
	// if err != nil {
	// 	t.Errorf("failed to drop index, got error %v\n", err)
	// }

	// Rename index failure
	// err := DB.Migrator().RenameIndex(&User{}, "name_idx", "name_idx2")
	// if err != nil {
	// 	t.Errorf("failed to rename index, got error %v\n", err)
	// }
}
