package main

import (
	"testing"

	"gorm.io/hints"
)

func TestImport(t *testing.T) {
	_ = hints.New("hint")
}
