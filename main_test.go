package main

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func where() string {
	return utils.FileWithLineNum()
}

func TestFileWithLineNumExternal(t *testing.T) {
	_, file, line, _ := runtime.Caller(0)
	expectedLine := line + 2
	got := where()
	expectedSuffix := fmt.Sprintf("%s:%d", filepath.ToSlash(file), expectedLine)
	require.Equal(t, expectedSuffix, got)
}

func TestSlogCallerSource(t *testing.T) {
	buf := &bytes.Buffer{}
	handler := slog.NewTextHandler(buf, &slog.HandlerOptions{AddSource: true})
	l := logger.NewSlogLogger(slog.New(handler), logger.Config{LogLevel: logger.Info})

	l.Trace(context.Background(), time.Now(), func() (string, int64) {
		return "select 1", 0
	}, nil)

	out := buf.String()
	require.NotContains(t, out, "gorm/logger/slog.go")
	require.Contains(t, out, "playground/main_test.go")
}
