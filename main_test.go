package main

import (
	"log"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	parent := Parent{
		ID:"1",
		Name: "name1",
		Children: []Child{
			{ID:"1", ParentID: "1", ChildName: "Child1"},
			{ID:"2", ParentID: "1", ChildName: "Child2"},
			{ID:"3", ParentID: "1", ChildName: "Child3"},
		},
	}

	/*
	DB.Create(&parent)
	*/
	err := DB.Create(&parent).Error
	log.Println(err)
}
