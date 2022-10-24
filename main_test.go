package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	company := &Company{Name: "ACME"}
	DB.Create(company)

	manager := &User{
		Name:      "manager",
		CompanyID: &company.ID,
		// have no manager
		ManagerID: nil,
	}
	DB.Create(manager)

	user := &User{
		Name:      "jinzhu",
		CompanyID: &company.ID,
		ManagerID: &manager.ID,
	}
	DB.Create(user)

	props := []*UserProp{
		{
			CompanyID: company.ID,
			ManagerID: manager.ID,
			Value:     "foo",
		},
	}
	DB.Create(&props)

	t.Run("user has props", testUser(user, "foo"))
	t.Run("user without props", testUser(manager, "bar"))
}

func testUser(user *User, propValue string) func(*testing.T) {
	return func(t *testing.T) {
		var resultUser User
		if err := DB.Preload("UserProps").First(&resultUser, user.ID).Error; err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
		if resultUser.UserProps == nil {
			t.Errorf("Failed, expected: %v, got %v", propValue, nil)
		} else if resultUser.UserProps.Value != propValue {
			t.Errorf("Failed, expected: %v, got %v", propValue, resultUser.UserProps.Value)
		}
	}
}
