package main

import (
	"database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"os"
	"regexp"
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
func TestGORMV2_Failing(t *testing.T) {

	type Student struct {
		gorm.Model
		Name string `gorm:"type:varchar(50);not null"`
	}
	type v2Suite struct {
		db      *gorm.DB
		mock    sqlmock.Sqlmock
		student Student
	}

	s := &v2Suite{}
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	s.db, err = gorm.Open(postgres.New(
		postgres.Config{
			Conn:       db,
			DriverName: "postgres",
		},
	), &gorm.Config{})
	if err != nil {
		panic(err) // Error here
	}

	defer db.Close()

	s.student = Student{
		Name: "Test 1",
	}

	defer db.Close()

	s.mock.ExpectBegin()

	s.mock.ExpectExec(
		regexp.QuoteMeta(`INSERT INTO "students" ("created_at","updated_at","deleted_at","name") VALUES ($1,$2,$3,$4) RETURNING "id"`)).
		WithArgs(AnyTime{}, AnyTime{}, nil, s.student.Name).
		WillReturnResult(sqlmock.NewResult(int64(1), 1))

	s.mock.ExpectCommit()

	if err = s.db.Create(&s.student).Error; err != nil {
		t.Errorf("Failed to insert to gorm db, got error: %v", err)
		t.FailNow()
	}

	err = s.mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}

func TestGORMV2_Working(t *testing.T) {
	type Student struct {
		ID   string `gorm:"primaryKey,autoIncrement"`
		Name string `gorm:"type:varchar(50);not null"`
	}
	type v2Suite struct {
		db      *gorm.DB
		mock    sqlmock.Sqlmock
		student Student
	}

	s := &v2Suite{}
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	s.db, err = gorm.Open(postgres.New(
		postgres.Config{
			Conn:       db,
			DriverName: "postgres",
		},
	), &gorm.Config{})
	if err != nil {
		panic(err) // Error here
	}

	defer db.Close()

	s.student = Student{
		Name: "Test 1",
	}

	defer db.Close()

	s.mock.ExpectBegin()

	s.mock.ExpectExec(
		regexp.QuoteMeta(`INSERT INTO "students" ("id","name") VALUES ($1,$2)`)).
		WithArgs("", s.student.Name).
		WillReturnResult(sqlmock.NewResult(int64(1), 1))

	s.mock.ExpectCommit()

	if err = s.db.Create(&s.student).Error; err != nil {
		t.Errorf("Failed to insert to gorm db, got error: %v", err)
		t.FailNow()
	}

	err = s.mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}

func TestGORMV2_ID_Working(t *testing.T) {

	type Student struct {
		gorm.Model
		Name string `gorm:"type:varchar(50);not null"`
	}
	type v2Suite struct {
		db      *gorm.DB
		mock    sqlmock.Sqlmock
		student Student
	}

	s := &v2Suite{}
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	s.db, err = gorm.Open(postgres.New(
		postgres.Config{
			Conn:       db,
			DriverName: "postgres",
		},
	), &gorm.Config{})
	if err != nil {
		panic(err) // Error here
	}

	defer db.Close()

	s.student = Student{
		Name: "Test 1",
	}

	defer db.Close()
	addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")

	s.mock.ExpectBegin()

	s.mock.ExpectQuery(`INSERT INTO "students" (.+) VALUES (.+) RETURNING "id"`).WillReturnRows(addRow)

	s.mock.ExpectCommit()

	if err = s.db.Create(&s.student).Error; err != nil {
		t.Errorf("Failed to insert to gorm db, got error: %v", err)
		t.FailNow()
	}

	err = s.mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}

func TestMain(m *testing.M) {

	code := m.Run()

	os.Exit(code)
}
