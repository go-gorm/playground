package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Role struct {
	ID     uint   `gorm:"primarykey"`
	Name   string `gorm:"unique;not null"`
	Entity string `gorm:"unique;not null"`
}

type Group struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"unique;not null"`
}

type RolesGroups struct {
	RID Role  `gorm:"foreignKey:id"`
	GID Group `gorm:"foreignKey:id"`
}

func TestGORM(t *testing.T) {
	err := DB.AutoMigrate(&Role{},
		&Group{},
		&RolesGroups{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
