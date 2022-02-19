package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func (gm *GroupMember) BeforeCreate(tx *gorm.DB) error {
	myValue, ok := tx.Get("role_value")
	// ok => true
	// myValue => 123
	if ok {
		gm.Role = myValue.(string)
	}

	return nil
}
func TestGORM(t *testing.T) {
	myValue := "myRole"
	user := User{Name: "jinzhu"}
	group := Group{Name: "group", Members: []*User{&user}}

	// Test set using create on the associated model
	DB.Set("role_value", myValue).Create(&group)

	var result GroupMember
	if err := DB.Where("user_id = ? AND group_id = ?", user.ID, group.ID).First(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result.Role != myValue {
		t.Errorf("Failed, groupMember.Role: '%s' want: '%s'", result.Role, myValue)
	}

	// Test set using association mode
	group2 := Group{Name: "group 2"}
	DB.Create(&group2)

	DB.Model(&group2).Set("role_value", myValue).Association("Members").Append(&user)

	var result2 GroupMember
	if err := DB.Where("user_id = ? AND group_id = ?", user.ID, group2.ID).First(&result2).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if result2.Role != myValue {
		t.Errorf("Failed, groupMember.Role: '%s' want: '%s'", result2.Role, myValue)
	}

}
