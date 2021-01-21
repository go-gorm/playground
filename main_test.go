package main

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
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

type Model struct {
	gorm.Model
}

func TestAutomigrate(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Errorf("sqlmock could not be created")
	}
	defer func() {
		if err := sqlMock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	}()

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: dbMock}), &gorm.Config{})
	if err != nil {
		t.Errorf("db could not be created")
	}

	// This test passing would be the expected behaviour
	t.Run("fail on create table", func(t *testing.T) {
		errMsg := "sql error"
		sqlMock.ExpectExec("CREATE TABLE").WillReturnError(errors.New(errMsg))
		// uncomment this will make the test pass
		//sqlMock.ExpectExec("CREATE INDEX").WillReturnError(errors.New(errMsg))

		err = db.AutoMigrate(&Model{})
		if err == nil {
			t.Errorf("expected migration to return error")
		}

		if err.Error() != errMsg {
			t.Errorf("expected migration failed with %s, but got %s", errMsg, err.Error())
		}
	})
}
