package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
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

func TestTimeZone(t *testing.T) {
	birthday := time.Date(2023, 11, 12, 9, 7, 18, 0, time.UTC)
	user := User{Name: "TimeZone UTC", Birthday: &birthday}

	if err := DB.Create(&user).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
		t.FailNow()
	}

	birthdayResult := *user.Birthday
	assert.Equal(t, birthday, birthdayResult)

	userResult := User{}

	if err := DB.Where("name = ?", "TimeZone UTC").First(&userResult).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
		t.FailNow()
	}

	birthdayResult = *userResult.Birthday
	assert.Equal(t, birthday, birthdayResult)

}

func TestPGXPooltoSqlDB(t *testing.T) {
	if DB.Dialector.Name() != "postgres" {
		t.Skip("test only runs in postgres")
	}

	// The following test shows how to open a pgxpool with the forced return of timezone as UTC.
	ctx := context.Background()

	uri := "user=gorm password=gorm host=localhost dbname=gorm port=9920 sslmode=disable TimeZone=UTC"

	config, err := pgxpool.ParseConfig(uri)
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
		t.FailNow()
	}

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		conn.TypeMap().RegisterType(&pgtype.Type{
			Name:  "timestamptz",
			OID:   pgtype.TimestamptzOID,
			Codec: &pgtype.TimestamptzCodec{ScanLocation: time.UTC}, // It would be better to parse the connection string for TimeZone and use that here.  Or have some other way of specifying it.
		})

		return nil
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
		t.FailNow()
	}

	// test connection to ensure all is well
	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		t.Errorf("Failed, got error: %v", err)
		t.FailNow()
	}

	nativeDB := stdlib.OpenDBFromPool(pool)
	db, err := gorm.Open(postgres.New(postgres.Config{Conn: nativeDB}))
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
		t.FailNow()
	}

	birthday := time.Date(2023, 11, 12, 9, 7, 18, 0, time.UTC)
	userResult := User{}

	if err := db.Where("name = ?", "TimeZone UTC").First(&userResult).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
		t.FailNow()
	}

	birthdayResult := *userResult.Birthday
	assert.Equal(t, birthday, birthdayResult)

	// Need some way to allow the closing of the connection pool after the gorm connection has closed.
	if sqlDB, _ := db.DB(); sqlDB != nil {
		sqlDB.Close()
		pool.Close()
	}
}
