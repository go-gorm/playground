package main

import (
	"testing"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/driver/sqlite"
)

type TestFile struct {
	ProjName string
	File string
}

type TestFileList struct {
	ProjName string `gorm:"primaryKey"`
	Files []TestFile `gorm:"ForeignKey:ProjName"`
}

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	var files []TestFile

	files = append(files,TestFile{"test","test"})

	zart := TestFileList {"test",files}

	_db,_ := gorm.Open(sqlite.Open("db.sql"), &gorm.Config{PrepareStmt: true})

	_db.AutoMigrate(&TestFile{},&TestFileList{})

	if err := _db.Clauses(clause.OnConflict{DoNothing: true}).Create(zart).Error ; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}