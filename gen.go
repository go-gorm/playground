package main

import (
	"gorm.io/gen"
	// "gorm.io/gen/examples/dal"
)

func generate() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/query",
		Mode:    gen.WithDefaultQuery, /*WithQueryInterface, WithoutContext*/
	})

	db, err := OpenTestConnection()
	if err != nil {
		panic(err)
	}

	RunMigrations()

	g.UseDB(db)
	g.ApplyBasic(g.GenerateModel("users"))
	g.Execute()
}
