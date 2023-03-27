package main

import (
	"gorm.io/gorm"
	"testing"
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

type Person struct {
	ID        string
	FirstName string
	LastName  string
}

func TestGORM_DoesNotInsertEmptyString_WhenPrimaryKeyIsZeroValue(t *testing.T) {
	// Arrange
	p := Person{
		FirstName: "John",
		LastName:  "Doe",
	}

	session := DB.Session(&gorm.Session{DryRun: true})
	var expectedSQL string
	switch DB.Dialector.Name() {
	case "postgres":
		expectedSQL = `INSERT INTO "people" ("first_name","last_name") VALUES ('John','Doe')`
	default:
		expectedSQL = "INSERT INTO `people` (`first_name`,`last_name`) VALUES ('John','Doe')"
	}

	// Act
	stmt := session.Create(&p).Statement
	actualSQL := DB.Dialector.Explain(stmt.SQL.String(), stmt.Vars...)

	// Assert
	if expectedSQL != actualSQL {
		t.Errorf(`Expected "%s" to equal "%s"`, expectedSQL, actualSQL)
	}
}
