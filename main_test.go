package main

import (
	db_sql "database/sql"
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestBindingValuesUsingRaw(t *testing.T) {
	dryRunDB := DB.Session(&gorm.Session{DryRun: true})

	query := &User{Name: "asdf", Age: 10}

	sqlStatement := dryRunDB.Where(query).Find(&User{}).Statement
	sqlStr := sqlStatement.SQL.String()

	rawRes := dryRunDB.Raw(sqlStr, "asdf", 10)
	boundSqlStr := rawRes.Statement.SQL.String()
	if boundSqlStr != "SELECT * FROM `users` WHERE `users`.`name` = \"asdf\" AND `users`.`age` = 10 AND `users`.`deleted_at` IS NULL" {
		t.Fatalf("invalid sql generated, got %v", boundSqlStr)
	}
}
func TestBindingValuesUsingRawWithNamedArg(t *testing.T) {
	dryRunDB := DB.Session(&gorm.Session{DryRun: true})

	query := &User{Name: "asdf", Age: 10}

	sqlStatement := dryRunDB.Where(query).Find(&User{}).Statement
	sqlStr := sqlStatement.SQL.String()

	nameArg := db_sql.Named("Name", "asdf")
	ageArg := db_sql.Named("Age", 10)
	rawRes := dryRunDB.Raw(sqlStr, nameArg, ageArg)
	boundSqlStr := rawRes.Statement.SQL.String()
	if boundSqlStr != "SELECT * FROM `users` WHERE `users`.`name` = \"asdf\" AND `users`.`age` = 10 AND `users`.`deleted_at` IS NULL" {
		t.Fatalf("invalid sql generated, got %v", boundSqlStr)
	}

}

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
