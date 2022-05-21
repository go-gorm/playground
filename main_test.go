package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql, sqlite

func TestGORM(t *testing.T) {
	DB.Exec("DROP TABLE IF EXISTS table_a")
	DB.Exec("CREATE TABLE table_a (a_name CHAR(8) NOT NULL)")

	var want string
	switch n := DB.Dialector.Name(); n {
	case "mysql":
		want = "INSERT INTO table_a (`a_name`) VALUES ('')"

	case "sqlite":
		want = "INSERT INTO `table_a` (`a_name`) VALUES (\"\")"

	default:
		t.Fatalf("Dialector not supported: %s", n)
	}

	sql := DB.ToSQL(fn)
	if sql != want {
		t.Fatalf("sql not match.\nwant: %s\ngot : %s", want, sql)
	}
	// sqlite : INSERT INTO `a` (`a_name`) VALUES ("")
	// mysql  : INSERT INTO table_a AS a (`a_name`) VALUES ('')

	// fn can't run directly.
	err := fn(DB).Error
	if err == nil {
		t.Fatal("err is nil")
	}
	// sqlite : no such table: a
	// mysql  : Error 1064: You have an error in your SQL syntax;
}

func fn(tx *gorm.DB) *gorm.DB {
	var data struct {
		AName string
	}
	tx = tx.Table("table_a AS a")
	tx = tx.Create(&data)
	return tx
}
