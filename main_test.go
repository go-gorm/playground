package main

import (
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

func main() {
	var files []TestFile

	files = append(files,TestFile{"test","test"})

	zart := TestFileList {"test",files}

	_db,_ := gorm.Open(sqlite.Open("db.sql"), &gorm.Config{PrepareStmt: true})

	_db.AutoMigrate(&TestFile{},&TestFileList{})

	_db.Clauses(clause.OnConflict{DoNothing: true}).Create(zart)
}