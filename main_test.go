package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	var result *gorm.DB

	// Try inserting into DB without manually specifying identity column value
	workTestStruct := TestStruct{Value: "hello"}
	result = DB.Create(&workTestStruct)
	if err := result.Error; err != nil {
		t.Errorf("Failed to create workTestStruct: %+v", err)
	}

	var ts1 TestStruct
	if err := DB.First(&ts1, workTestStruct.ID).Error; err != nil {
		t.Errorf("Failed to read workTestStruct, got error: %v", err)
	}

	// Try again but specifying the identity column value.
	// The insert should fail here due to the schema not being included in the `SET IDENTITY_INSERT` statement
	failTestStruct := TestStruct{ID: 100, Value: "there"}
	result = DB.Create(&failTestStruct)
	if err := result.Error; err != nil {
		t.Errorf("Failed to create failTestStruct: %+v", err)
	}

	var ts2 TestStruct
	if err := DB.First(&ts2, failTestStruct.ID).Error; err != nil {
		t.Errorf("Failed to read failTestStruct, got error: %v", err)
	}
}
