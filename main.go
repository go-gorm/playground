package main

import (
	"fmt"
	"os"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

type Test struct {
	ID            int
	Value         int
	OptionalValue *decimal.Decimal
}

func main() {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("gorm_%d.db", time.Now().Unix())), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.Exec(`CREATE TABLE tests (
		id int not null,
		value int,
		optional_value int
	)`).Error; err != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {
		value := interface{}(nil)
		if i%2 == 0 {
			value = i
		}
		db.Exec(`INSERT INTO tests (id, value, optional_value) VALUES (?, ?, ?)`, i, value, value)
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

			if test.OptionalValue == nil {
				fmt.Printf("%s:MISSMATCH: id=%d optional_value=null expected_value=%d\n", os.Args[1], test.ID, test.ID)
				nrError++
			} else {
				if test.ID != int(test.OptionalValue.InexactFloat64()) {
					fmt.Printf("%s:MISSMATCH: id=%d optional_value=%d expected_value=%d\n", os.Args[1], test.ID, int(test.OptionalValue.InexactFloat64()), test.ID)
					nrError++
				}
			}

		} else {
			if test.Value != 0 {
				fmt.Printf("%s:MISSMATCH: id=%d value=%d expected_value=0\n", os.Args[1], test.ID, test.Value)
				nrError++
			}

			if test.OptionalValue != nil {
				fmt.Printf("%s:MISSMATCH: id=%d optional_value=%f expected_value=0\n", os.Args[1], test.ID, test.OptionalValue.InexactFloat64())
				nrError++
			}
		}
	}

	if nrError == 0 {
		fmt.Printf("%s:PASS\n", os.Args[1])
	}

}
