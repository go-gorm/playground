package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type M struct {
	ID uint64 `gorm:"primarykey"`
	I  string `gorm:"uniqueIndex:idx_name"`
	I2 string `gorm:"uniqueIndex:idx_name"`
}

func Test(t *testing.T) {
	r := require.New(t)
	db, err := gorm.Open(sqlite.Open("test.sqlite3?_journal=WAL"), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
			},
		),
	})
	r.NoError(err)
	db.AutoMigrate(&M{})
}
