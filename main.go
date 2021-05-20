package main

import (
	"fmt"
	"gorm.io/gorm/clause"
)

type Base struct {
	ID int `gorm:"primaryKey;unique"`
}
type PackageName struct {
	Base
	Name string `gorm:"unique"`
}

var onConflict = clause.OnConflict{
	Columns: []clause.Column{{Name: "name"}},
	DoUpdates: []clause.Assignment{{
		Column: clause.Column{Name: "name"},
		Value:  clause.Column{Table: "excluded", Name: "name"},
	}},
}

func main() {
	var written []PackageName
	var read []PackageName

	for i := 0; i < 10; i++ {
		written = append(written, PackageName{Name: fmt.Sprintf("N%v", i+1)})
	}
	// Create items
	DB.Clauses(onConflict).CreateInBatches(&written, 5)
	fmt.Printf("Returned: \n")
	for _, n := range written {
		fmt.Printf("id:%d name:%v\n", n.ID, n.Name)
	}

	// Do something in the middle
	DB.Create(&User{})

	// Create a set of new items, partly overlapping
	for i := 0; i < 10; i++ {
		written[i].ID = 0
		written[i].Name = fmt.Sprintf("N%d", 5+i)
	}

	// Write a second set of items
	DB.Clauses(onConflict).CreateInBatches(&written, 5)
	fmt.Printf("Returned second: \n")
	for _, n := range written {
		fmt.Printf("id:%d name:%v\n", n.ID, n.Name)
	}

	// Read back actual database state
	DB.Find(&read)
	fmt.Printf("Actual: \n")
	for _, n := range read {
		fmt.Printf("id:%d name:%v\n", n.ID, n.Name)
	}
}
