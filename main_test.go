package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type MyTable struct {
	Def string `gorm:"size:512;index:idx_def,unique"`
	Abc string `gorm:"size:65000000"`
}

func TestGORM(t *testing.T) {
	sql := "CREATE TABLE `my_tables` (`def` varchar(512),`abc` longtext,UNIQUE INDEX `idx_def` (`def`))"
	if err := DB.Exec(sql).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	// stdout to pipe
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	DB.DryRun = true
	if err := DB.AutoMigrate(&MyTable{}); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// get stdout from pipe
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if strings.Contains(string(out), "ALTER TABLE") {
		t.Errorf("Unexpected alter table: %v", string(out))
	}
}
