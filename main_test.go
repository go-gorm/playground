package main

import (
	"context"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gen"
	"gorm.io/playground/dal/query"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
func TestGen(t *testing.T) {
	RunMigrations()
	generate()
	query.SetDefault(DB)
	list, err := query.User.WithContext(context.Background()).Where(gen.Cond(datatypes.JSONArrayQuery("area").Contains([]int{1, 2, 3}))...).Find()
	if err != nil {
		// SELECT * FROM `users` WHERE JSON_CONTAINS(`area`,JSON_ARRAY((1,2,3))) AND `users`.`deleted_at` IS NULL
		// Error 1241 (21000): Operand should contain 1 column(s)
		t.Errorf("Failed, got error: %v", err)
	} else {
		fmt.Printf("%#v\n", len(list))
	}
}
