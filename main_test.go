package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

func TestGORM(t *testing.T) {

	nonExistingFolderId := 65473574
	folder := Folder{
		Name:     "I reference non existing",
		ParentID: &nonExistingFolderId,
	}

	result := DB.Save(&folder)
	if result.Error == nil {
		t.Errorf("expected foreign key violation, but got no error")
	}
}
