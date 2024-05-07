package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver
func TestGORM(t *testing.T) {
	user := User{
		Name: "jinzhu",
		Account: Account{
			Number: "123456",
			Companies: []Company{
				{Name: "Corp1"}, {Name: "Corp2"},
			},
			Pet: Pet{
				Name: "Pet1",
			},
		},
	}
	DB.Create(&user)

	for i := 0; i < 20; i++ {
		DB.Create(&User{Name: fmt.Sprintf("User%d", i)})
	}

	var entries []User
	assert.NotPanics(t, func() {
		_ = DB.WithContext(context.Background()).
			Joins("Account.Pet").
			Preload("Account.Companies").
			Find(&entries).Error
	})

}
