package main

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

type Test struct {
	ID    int
	Value int
}

func main() {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("gorm_%d.db", time.Now().Unix())), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.Exec(`CREATE TABLE tests (
		id int not null,
		value int
	)`).Error; err != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {
		value := interface{}(nil)
		if i%2 == 0 {
			value = i
		}
		db.Exec(`INSERT INTO tests (id, value) VALUES (?, ?)`, i, value)
	}

	// let's read the whole thing

	rows, err := db.Model(&Test{}).Rows()
	if err != nil {
		panic(err)
	}

	test := Test{}

	nrError := 0
	defer rows.Close()
	for rows.Next() {

		if err := db.ScanRows(rows, &test); err != nil {
			panic(err)
		}

		if test.ID%2 == 0 {
			if test.ID != test.Value {
				fmt.Printf("%s:MISSMATCH: id=%d value=%d expected_value=%d\n", os.Args[1], test.ID, test.Value, test.ID)
				nrError++
			}
		} else {
			if test.Value != 0 {
				fmt.Printf("%s:MISSMATCH: id=%d value=%d expected_value=0\n", os.Args[1], test.ID, test.Value)
				nrError++
			}
		}
	}

	if nrError == 0 {
		fmt.Printf("%s:PASS\n", os.Args[1])
	}

}
