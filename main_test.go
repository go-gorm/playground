package main

import (
	"strings"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Gateway struct {
	IP  string `gorm:"index:uniq_vip,unique"`
	UIN string `gorm:"index:uniq_vip,unique"`
}

func TestGORM(t *testing.T) {
	readLog() // clean history log

	// create not exist table first time
	err := DB.AutoMigrate(Gateway{})
	if err != nil {
		t.Error(err)
	}

	readLog()

	// should do noting
	err = DB.AutoMigrate(Gateway{})
	if err != nil {
		t.Error(err)
	}

	migrateLog := readLog() // read this migration log

	if strings.Contains(migrateLog, "CREATE") || strings.Contains(migrateLog, "DROP") {
		t.Errorf("second migration should not do recreate table")
		t.Errorf("migration log: %s", migrateLog)
		t.FailNow()
	}

}
