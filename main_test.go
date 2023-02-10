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

	props := []*UserProp{
		{
			CompanyID: company.ID,
			ManagerID: manager.ID,
			Value:     "foo",
		},
	}
	DB.Create(&props)

	user := &User{
		Name:      "jinzhu",
		CompanyID: &company.ID,
		ManagerID: &manager.ID,
	}
	DB.Create(user)

	value := "foo"
	t.Run("user has props", testUser(user.ID, &value))
	t.Run("user without props", testUser(manager.ID, nil))
}

func testUser(userID uint, expectedValue *string) func(*testing.T) {
	return func(t *testing.T) {
		var actualUser User

		err := DB.Preload("UserProps").First(&actualUser, userID).Error
		if err != nil {
			t.Fatalf("Failed, got error: %v", err)
		}

		if actualUser.UserProps == nil {
			if expectedValue != nil {
				t.Fatalf("Failed, expected: %v, got %v", *expectedValue, nil)
			}
		}

		if expectedValue != nil {
			if actualUser.UserProps.Value != *expectedValue {
				t.Fatalf("Failed, expected: %v, got %v", *expectedValue, actualUser.UserProps.Value)
			}
		}

		if actualUser.UserProps != nil && expectedValue == nil {
			t.Fatalf("Failed, expected: %v, got %v", nil, actualUser.UserProps.Value)
		}
	}
}
