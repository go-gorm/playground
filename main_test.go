package main

import (
	"os"
	"strings"
	"testing"

	"ariga.io/atlas-provider-gorm/gormschema"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestDBMigrationGeneration(t *testing.T) {
	dialect := os.Getenv("GORM_DIALECT")
	stmts, err := gormschema.New(dialect).Load(&Account{})

	if err != nil {
		t.Errorf("error loading gorm schema: %v", err)
	}

	if !strings.Contains(stmts, "CREATE TABLE") {
		t.Errorf(`"CREATE TABLE" not found in generated statements: %s`, stmts)
	}
}
