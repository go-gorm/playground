package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	row := NoDefaultID{Name: "test"}
	row2 := NoDefaultID{Name: "test2"}

	DB.Create(&row).Scan(&row)
	DB.Create(&row2).Scan(&row2)

	fmt.Println(row.ID)
	fmt.Println(row2.ID)
	if row.ID != row2.ID {
		t.Error("ID shouldn't be auto incremented")
	}
}
