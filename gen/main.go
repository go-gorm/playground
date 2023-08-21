package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/playground/models"
)

func main() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"

	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query", // output directory, default value is ./query
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// Initialize a *gorm.DB instance
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	g.UseDB(db)
	g.ApplyBasic(models.User{})
	g.Execute()
}
