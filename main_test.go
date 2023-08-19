package main

import (
	"testing"

	"gorm.io/gen"
	"gorm.io/playground/query"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql

func TestGORM_GEN(t *testing.T) {
	config := gen.Config{
		OutPath:          "./query",
		FieldWithTypeTag: true,
		FieldNullable:    true,
		Mode:             gen.WithDefaultQuery | gen.WithQueryInterface,
	}
	g := gen.NewGenerator(config)
	g.UseDB(DB)
	g.ApplyInterface(
		func(query.CustomQuery) {},
		g.GenerateModel("users"),
	)
	g.Execute()
}
