package main

import (
	"strconv"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	positiveInf, _ := strconv.ParseFloat("+Inf", 64)
	negativeInf, _ := strconv.ParseFloat("-Inf", 64)

	testCases := []User{
		{Point: 0.0},         // Will pass
		{Point: positiveInf}, // Fail
		{Point: negativeInf}, // Fail
	}

	for _, tc := range testCases {
		if err := DB.Create(&tc).Error; err != nil {
			t.Errorf("Failed, got error: %v", err)
		}

		var result User
		if err := DB.First(&result, tc.ID).Error; err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
	}
}
