package main

import (
	"gorm.io/gorm/clause"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

/*
 PASS
*/
func TestModelWithoutIgnoredPro(t *testing.T) {
	tom := PersonOne{
		ID: 1,
		Friends: []PersonOne{
			{
				ID: 2,
				Friends: []PersonOne{
					{
						ID: 3,
						Friends: []PersonOne{
							{
								ID: 4,
							},
						},
					},
				},
			},
		},
	}
	DB.Create(&tom)

	var result PersonOne
	DB.Preload("Friends.Friends.Friends."+clause.Associations).Find(&result, "id = ?", 1)
	if !PersonOneEqual(tom, result) {
		t.Errorf("Failed, \"tom\" is not created correctly in sql")
	}
}

/*
 FAIL
 You can debug to see the details about "tom" and "result"
 Compare with previous TestModelWithoutIgnoredPro
*/
func TestModelWithIgnoredPro(t *testing.T) {
	tom := PersonTwo{
		ID: 1,
		Friends: []PersonTwo{
			{
				ID: 2,
				Friends: []PersonTwo{
					{
						ID: 3,
						Friends: []PersonTwo{
							{
								ID: 4,
							},
						},
					},
				},
			},
		},
	}
	DB.Create(&tom)

	var result PersonTwo
	DB.Preload("Friends.Friends.Friends."+clause.Associations).Find(&result, "id = ?", 1)
	if !PersonTwoEqual(tom, result) {
		t.Errorf("Failed, \"tom\" is not created correctly in sql")
	}
}
