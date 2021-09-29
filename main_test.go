package main

import (
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
	"gotest.tools/assert"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestCreateUserWithNestedAssociations(t *testing.T) {
	user := User{
		Name: "jinzhu",
		Friends: []*User{
			{
				Name: "gorm",
				Languages: []Language{
					{
						Code: "EN",
						Name: "English",
					},
				},
			},
		},
	}

	DB.Create(&user)

	var result User
	if err := DB.Preload("Friends.Languages").First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	assert.DeepEqual(t, result, user, cmpopts.IgnoreFields(User{}, "Model"))
}

func TestUpdateUserWithNestedAssociations(t *testing.T) {
	user := User{
		Name: "john",
		Friends: []*User{
			{
				Name: "doe",
				Languages: []Language{
					{
						Code: "EN",
						Name: "English",
					},
				},
			},
		},
	}

	DB.Create(&user)

	var result User
	if err := DB.Preload("Friends.Languages").First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	assert.DeepEqual(t, result, user, cmpopts.IgnoreFields(User{}, "Model"))

	user.Friends = []*User{
		{
			Name: "paul",
			Languages: []Language{
				{
					Code: "FR",
					Name: "French",
				},
			},
		},
	}

	var updated User
	if err := DB.Preload("Friends.Languages").Updates(user).Scan(&updated).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	assert.DeepEqual(t, updated, user, cmpopts.IgnoreFields(User{}, "Model"))
}
