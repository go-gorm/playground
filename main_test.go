package main

import (
	"fmt"
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Foo struct {
	ID        string `gorm:"primaryKey;"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func TestGORM(t *testing.T) {
	DB.AutoMigrate(&Foo{})

	f := Foo{
		ID:   "id",
		Name: "jinzhu",
	}

	if err := DB.Create(&f).Error; err != nil {
		panic(err)
	}

	fmt.Println(f)
	f2 := Foo{
		ID:   "id",
		Name: "new-name",
	}

	if err := DB.Save(&f2).Error; err != nil {
		panic(err)
	}

	fmt.Println(f2)

}
