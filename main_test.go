package main

import (
	"testing"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func populate(db *gorm.DB) error {
	parents := []*Parent{
		{
			ID: 5,
		},
		{
			ID: 6,
		},
		{
			ID: 7,
		},
		{
			ID: 8,
		},
	}
	return db.Create(&parents).Error

}


func TestGORM(t *testing.T) {
	err := populate(DB)
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	preparedDB := DB.Where("id <> 0")
	var parents []*Parent
	if err = preparedDB.Where("id in (?)", []int{5, 6}).Find(&parents).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if len(parents) != 2 {
		t.Errorf("Failed, Expected 2 parents.  Got %d", len(parents))
	}
	if err = preparedDB.Where("id in (?)", []int{7, 8}).Find(&parents).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if len(parents) != 2 {
		t.Errorf("Failed, Expected 2 parents.  Got %d", len(parents))
	}
}
