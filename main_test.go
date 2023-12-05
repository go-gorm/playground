package main

import (
	"bytes"
	"io"
	"log"
	"strings"
	"testing"
	"time"

	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	birthday := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	user := User{Name: "jinzhu", Birthday: &birthday}

	DB.Create(&user)

	var buf bytes.Buffer
	DB.Logger = logger.New(log.New(io.Writer(&buf), "", 0), logger.Config{
		LogLevel: logger.Info,
	})

	DB.Use(dbresolver.Register(dbresolver.Config{
		TraceResolverMode: true,
	}))

	var result User
	if err := DB.Where("birthday = ?", birthday).First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	substr := "[1990-01-01 00:00:00 +0000 UTC 1]"
	if strings.Contains(buf.String(), substr) {
		t.Errorf("Should not contain %v, but got %v", substr, buf.String())
	}
}
