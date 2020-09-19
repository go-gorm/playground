package main

import (
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Person struct {
	ID       string    `gorm:"column:id"`
	Name     string    `gorm:"column:name"`
	CreateAt time.Time `gorm:"column:create_at"`
}

func TestGORM(t *testing.T) {
	if err := DB.AutoMigrate(&Person{}); err != nil {
		t.Fail()
	}

	DB.Create(&Person{
		ID:       "1001",
		Name:     "Alice",
		CreateAt: time.Unix(1600480219, 0),
	})

	DB.Create(&Person{
		ID:       "1002",
		Name:     "Bob",
		CreateAt: time.Unix(1600480220, 0),
	})

	DB.Create(&Person{
		ID:       "1003",
		Name:     "Cathy",
		CreateAt: time.Unix(1600480221, 0),
	})

	var persons []Person
	var count int64

	res := DB.Limit(2).Offset(2).Order("create_at desc").Find(&persons).Limit(-1).Offset(-1).Count(&count)

	if res.Error != nil {
		t.Fail()
	}

	t.Log(persons) // [{1003 Cathy 2020-09-19 09:50:21 +0800 CST}]
	t.Log(count) // 3
}
