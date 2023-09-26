package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	//"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/playground/db/model"
)

func main() {
	conn := "host=localhost user=root password=secret dbname=simple_bank port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	fmt.Println(err)
	if err == nil {

		/* the follwoing code is used to generate the structs from the Postgres database */
		// g := gen.NewGenerator(gen.Config{
		// 	OutPath: "./db/queries",
		// 	Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		// })
		// g.UseDB(db) // reuse your gorm db
		// g.GenerateAllTable()
		// g.Execute()
		user := model.Account{AccountOwner: "Jinzhu", Balance: 1800, Currency: "USD"}

		result := db.Create(&user) // pass pointer of data to Create

		fmt.Println(user.AccountID, result.Error)
	}
}
