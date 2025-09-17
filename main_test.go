package main

import (
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

// TestGORM_FirstOrCreate tests the FirstOrCreate method
// this test demonstrates that when using gorm.Expr, the value is not updated in the struct UNLESS you use the Returning clause
// this behaviour should be fixed or at least documented
func TestGORM_FirstOrCreate(t *testing.T) {
	var result models.User
	if err := DB.Attrs(models.User{Name: "jinzhu", Age: 18}).FirstOrCreate(&result).Error; err != nil {
		t.Fatalf("Failed, got error: %v", err)
	} else {
		if result.Name != "jinzhu" {
			t.Errorf("Failed, got user name: %v, expected: %v", result.Name, "jinzhu")
		} else {
			t.Logf("Success, got user name: %v, expected: %v", result.Name, "jinzhu")
		}
	}

	// setting age to 19 and expect 19 in the struct, this works
	var update19 models.User
	age19 := 19
	if err := DB.Where("name = ?", "jinzhu").Assign("age", age19).FirstOrCreate(&update19).Error; err != nil {
		t.Fatalf("Failed, got error: %v", err)
	} else {
		if update19.Age != 19 {
			t.Errorf("Failed, got user age: %v, expected: %v", update19.Age, 19)
		} else {
			t.Logf("Success, got user age: %v, expected: %v", update19.Age, 19)
		}
	}

	// add 1 to age and expect 20 in the struct... but it fails
	var update20 models.User
	var age20 uint = 20
	if err := DB.Where("name = ?", "jinzhu").Assign("age", gorm.Expr("age + ?", 1)).FirstOrCreate(&update20).Error; err != nil {
		t.Fatalf("Failed, got error: %v", err)
	} else {
		if update20.Age != age20 {
			t.Errorf("Failed, got user age: %v, expected: %v", update20.Age, age20)
		} else {
			t.Logf("Success, got user age: %v, expected: %v", update20.Age, age20)
		}
	}

	// check if age is 20 in the db... it is!? why did it fail above?
	var selectUpdate20 models.User
	if err := DB.Where("name = ?", "jinzhu").First(&selectUpdate20).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	} else {
		if selectUpdate20.Age != 20 {
			t.Errorf("Failed, got user age: %v, expected: %v", selectUpdate20.Age, 20)
		} else {
			t.Logf("Success, got user %d with age %v", selectUpdate20.ID, selectUpdate20.Age)
		}
	}

	// okay, lets throw a returning in there and... it worked!? lmao
	var update2 models.User
	if res := DB.Clauses(clause.Returning{}).Where("name = ?", "jinzhu").Assign("age", gorm.Expr("age + ?", 5)).FirstOrCreate(&update2); res.Error != nil {
		t.Fatalf("Failed, got error: %v", res.Error)
	} else {
		if update2.Age != 25 {
			t.Errorf("Failed, got user age: %v, expected: %v", update2.Age, 25)
		} else {
			t.Logf("Success, got user age: %v, expected: %v", update2.Age, 25)
		}
	}
}
