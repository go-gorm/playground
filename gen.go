package main

import (
	"gorm.io/gen"
	"gorm.io/gen/examples/dal"
	"playground/model"
)

func generate() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/query",
		Mode:    gen.WithDefaultQuery, /*WithQueryInterface, WithoutContext*/

		WithUnitTest: true,
	})
	g.UseDB(dal.DB)

	g.ApplyBasic(model.User{})
	g.ApplyBasic(model.UserExt{})
	g.ApplyBasic(model.UserAccountRelation{})
	g.ApplyBasic(model.Account{})
	g.ApplyBasic(model.Company{})

	g.Execute()
}
