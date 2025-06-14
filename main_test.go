package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	man := Man{Id: 1, Name: "Man1", Age: 1}
	DB.Create(&man)

	// update data
	idx := Man{Id: 1}

	change1 := Man{
		Age: 10,
	}
	change2 := map[string]interface{}{
		"age": 20,
	}
	change3 := struct {
		Age int
	}{Age: 30}

	_, _, _ = change1, change2, change3
	var err error

	// change1 -> yes
	err = idx.update(change1)
	if err != nil {
		return
	}

	// change2 -> yes
	err = idx.update(change2)
	if err != nil {
		return
	}

	// change3 -> panic: reflect: Field index out of range
	err = idx.update(change3)
	if err != nil {
		return
	}
}
