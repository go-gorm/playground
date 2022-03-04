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
		ID:1,
		Name: "name1",
		Children: []Child{
			{ChildID:2, ParentID: 1, ChildName: "Child1"},
			{ChildID:3, ParentID: 1, ChildName: "Child2"},
			{ChildID:4, ParentID: 1, ChildName: "Child3"},
		},
	}

	parent2 := Parent{
		ID:2,
		Name: "name2",
		Children: []Child{
			{ChildID:5, ParentID: 2, ChildName: "Child4"},
			{ChildID:6, ParentID: 2, ChildName: "Child5"},
			{ChildID:6, ParentID: 2, ChildName: "Child6"},
		},
	}

	DB.Create(&parent)
	DB.Create(&parent2)

	var result Parent
	if err := DB.Model(&result).Select("*").Joins("LEFT JOIN children ON children.parent_id = parents.id").Where("parents.id = ?", 1).Scan(&result).Scan(&result.Children).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	log.Println(result)
}
