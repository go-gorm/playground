package main

import (
	"testing"

	"gorm.io/gorm"
)

func TestCompileError(t *testing.T) {
	dbg := DB.Session(&gorm.Session{DryRun: true})
	db := gorm.G[gorm.Model](dbg)

	var _ gorm.ChainInterface[gorm.Model] = db
	var _ gorm.Interface[gorm.Model] = db
}
