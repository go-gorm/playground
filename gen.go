package main

import (
	"gorm.io/gen"
	"gorm.io/gen/examples/dal"
)

func generate() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/query",
		Mode:    gen.WithDefaultQuery, /*WithQueryInterface, WithoutContext*/

		WithUnitTest: true,
	})
	g.UseDB(dal.DB)

	g.ApplyBasic(Company{}, Language{}) // Associations
	g.ApplyBasic(g.GenerateModel("user"), g.GenerateModelAs("account", "AccountInfo"))

	g.Execute()
}
