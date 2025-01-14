package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := A{Name: "jinzhu"}

	DB.Create(&user)

	type AB struct {
		A
		*B
	}
	a, b, c := &A{}, &B{}, &C{}
	aTable, bTable, cTable := a.TableName(), b.TableName(), c.TableName()
	var result1 AB
	err := DB.Model(a).Select(aTable+".*, "+bTable+".*").Joins(
		"LEFT JOIN "+bTable+" ON "+aTable+".id = "+bTable+".a_id",
	).Where(aTable+".id = ?", user.ID).Scan(&result1).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	fmt.Println("This result is in line with expectations: ")
	fmt.Println("no match b and b is nil: ", result1)

	type AC struct {
		A
		*C
	}
	var result2 AC
	err = DB.Model(a).Select(aTable+".*, "+cTable+".*").Joins(
		"LEFT JOIN "+cTable+" ON "+aTable+".id = "+cTable+".a_id",
	).Where(aTable+".id = ?", user.ID).Scan(&result2).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	fmt.Println("This result is not as expected: ")
	fmt.Println("no match c but c is not nil: ", result2)
}
