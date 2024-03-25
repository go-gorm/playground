package main

import (
	"gorm.io/gen"
	"gorm.io/playground/model"
	// "gorm.io/gen/examples/dal"
)

func generate() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/query",
		Mode:    gen.WithDefaultQuery, /*WithQueryInterface, WithoutContext*/

		WithUnitTest: true,
	})
	// g.UseDB(dal.DB)

	g.ApplyBasic(
		model.Company{},
		model.Language{},
		model.Account{},
		model.Pet{},
		model.Toy{},
		model.User{},
	) // Associations
	// g.ApplyBasic(g.GenerateModel("user"), g.GenerateModelAs("account", "AccountInfo"))

	g.Execute()
}
