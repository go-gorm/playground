package main

import (
	"bytes"
	"context"
	"log/slog"
	"strings"
	"testing"
	"time"

	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := models.User{Name: "jinzhu"}

	DB.Create(&user)

	var result models.User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestSlogLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	handler := slog.NewTextHandler(buf, &slog.HandlerOptions{AddSource: true})
	logger := NewSlogLogger(slog.New(handler), Config{LogLevel: Info})

	logger.Trace(context.Background(), time.Now(), func() (string, int64) {
		return "select count(*) from users", 0
	}, nil)

	if strings.Contains(buf.String(), "gorm/logger/slog.go") {
		t.Error("Found internal slog.go reference in caller frame. Expected only test file references.")
	}

	if !strings.Contains(buf.String(), "main_test.go") {
		t.Error("Missing expected test file reference. 'main_test.go' should appear in caller frames.")
	}
}

// func TestGORMGen(t *testing.T) {
// 	user := models.User{Name: "jinzhu2"}
// 	ctx := context.Background()

// 	gorm.G[models.User](DB).Create(ctx, &user)

// 	if u, err := gorm.G[models.User](DB).Where(g.User.ID.Eq(user.ID)).First(ctx); err != nil {
// 		t.Errorf("Failed, got error: %v", err)
// 	} else if u.Name != user.Name {
// 		t.Errorf("Failed, got user name: %v", u.Name)
// 	}
// }
