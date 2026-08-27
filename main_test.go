package main

import (
	"bytes"
	"log"
	"strings"
	"testing"

	"gorm.io/gorm/logger"
	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

// logger.Config.ParameterizedQueries is honoured for statements executed through
// the callback pipeline, but ignored by DB.Scan, which logs the statement with
// parameter values interpolated.
//
// DB.Scan replaces the connection's logger with logger.Recorder for the inner
// execution. That recorder is seeded at package level from logger.Default, so the
// connection's logger — and therefore its Config — is never consulted, and the
// values are interpolated into the recorded SQL before the real logger sees it.
func TestGORM(t *testing.T) {
	const secret = "SHOULD_NOT_BE_LOGGED"

	var buf bytes.Buffer
	DB.Logger = logger.New(log.New(&buf, "", 0), logger.Config{
		LogLevel:             logger.Info,
		ParameterizedQueries: true,
	})

	assertNotLogged := func(t *testing.T, what string) {
		t.Helper()
		if strings.Contains(buf.String(), secret) {
			t.Errorf("%s logged the parameter value despite ParameterizedQueries=true:\n%s",
				what, strings.TrimSpace(buf.String()))
		}
	}

	// These three keep the placeholder, as expected.
	buf.Reset()
	var users []models.User
	DB.Where("name = ?", secret).Find(&users)
	assertNotLogged(t, "Find")

	buf.Reset()
	var names []string
	DB.Model(&models.User{}).Where("name = ?", secret).Pluck("name", &names)
	assertNotLogged(t, "Pluck")

	buf.Reset()
	DB.Raw("SELECT name FROM users WHERE name = ?", secret).Row()
	assertNotLogged(t, "Raw+Row")

	// Scan interpolates it. Not raw-specific: the Raw query above is fine, while
	// this plain model query is not.
	buf.Reset()
	var scanned []models.User
	DB.Model(&models.User{}).Where("name = ?", secret).Scan(&scanned)
	assertNotLogged(t, "Scan")
}
