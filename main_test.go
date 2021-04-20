package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	tx := DB.Begin()
	defer tx.Rollback()

	if err := tx.Exec(
		`INSERT INTO "users" ( "id" ) VALUES ( $1 )`,
		`1`,
	).Error; err != nil {
		t.Errorf("Failed, got error:%v", err)
	}
}
