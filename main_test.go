package main

import (
	"fmt"
	"strings"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	DB.Create(&User{Name: "jinzhu1"})
	DB.Create(&User{Name: "jinzhu2"})
	DB.Create(&User{Name: "jinzhu3"})

	users := []User{}

	// this one works
	err := DB.Debug().
		Model(&User{}).
		Joins(fmt.Sprintf("JOIN unnest(ARRAY%s::int[]) WITH ORDINALITY AS x(id, order_nr) ON x.id = users.id", strings.ReplaceAll(fmt.Sprintf("%v", []uint{2, 1, 3}), " ", ","))).
		Order("x.order_nr").
		Find(&users).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// this one fails
	err = DB.Debug().
		Model(&User{}).
		Joins("JOIN unnest(ARRAY[?]::int[]) WITH ORDINALITY AS x(id, order_nr) ON x.id = users.id", []uint{2, 1, 3}).
		Order("x.order_nr").
		Find(&users).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
