package main

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func TestErrors(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Unable to create the mock db: %v", err)
		return
	}
	defer func() { _ = dbMock.Close() }()

	db, err := gorm.Open(mysql.Dialector{
		Config: &mysql.Config{
			Conn:                      dbMock,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{
		Logger: DB.Logger,
	})
	if err != nil {
		t.Errorf("Unable to create the gorm connection to the mock: %v", err)
		return
	}

	expectedError := errors.New("some error")
	mock.ExpectQuery("^SELECT ").WillReturnError(expectedError)

	var result []map[string]interface{}
	err = db.Table("users").Scan(&result).Error
	if err != expectedError {
		t.Errorf("Got wrong error: %v", err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %v", err)
	}
}
