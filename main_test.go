package main

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

func TestGORM(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	}), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	const id = 1
	const name = "b"
	var out *User
	gdb.Model(&User{}).Where("active", true).Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where(
			d.Where("id", id).Or("name", name),
		)
	}).First(&out)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "active" = $1 AND ("id" = $2 or "name" = $3) AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(true, id, name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(id))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
