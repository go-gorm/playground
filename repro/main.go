package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(postgres.Open(os.Getenv("GORM_DSN")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Run only against a dedicated disposable database.
	if err := db.Exec("DROP TABLE IF EXISTS users").Error; err != nil {
		log.Fatal(err)
	}
	if err := db.Exec("CREATE TABLE users (id bigserial PRIMARY KEY, name text, age bigint)").Error; err != nil {
		log.Fatal(err)
	}
	rows := []map[string]interface{}{{"Name": "alice", "Age": 18}}
	result := db.Debug().Model(&User{}).CreateInBatches(rows, 500)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	var users []User
	if err := db.Find(&users).Error; err != nil {
		log.Fatal(err)
	}
	if result.RowsAffected != 1 || len(users) != 1 || users[0].Name != "alice" || users[0].Age != 18 {
		log.Fatalf("unexpected result: rows=%d users=%+v", result.RowsAffected, users)
	}
	log.Printf("success: rows=%d users=%+v input=%v", result.RowsAffected, users, rows)
}
