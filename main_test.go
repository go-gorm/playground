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
	}
	err := db.Create(&parents).Error
	if err != nil {
		return err
	}
	children := []*Child{
		{
			ID:       1,
			ParentID: 5,
		},
		{
			ID:       2,
			ParentID: 5,
		},
	}
	return db.Create(&children).Error
}


func TestGORM(t *testing.T) {
	err := populate(DB)
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	var parent Parent
	err = DB.Preload("Children").Take(&parent, "id = 5").Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	db := DB.Unscoped()
	for _, c := range parent.Children {
		res := db.Delete(&Child{}, "id = ?", c.ID)
		if err = res.Error; err != nil {
		        t.Errorf("Failed, got error: %v", err)
		}
		if res.RowsAffected != 1 {
		        t.Errorf("Failed, expected 1 affected row for child %d. Got %d.", c.ID, res.RowsAffected)
		}
	}
	if err = db.Delete(&Parent{}, "id = ?", parent.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
