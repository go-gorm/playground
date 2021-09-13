package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite

func TestGORM(t *testing.T) {
	userA := User{}
	userB := User{}

	chat := Chat{}

	DB.Create(&userA)
	DB.Create(&userB)
	DB.Create(&chat)

	var resultA User
	if err := DB.First(&resultA, userA.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var resultB User
	if err := DB.First(&resultB, userB.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// userA has userB blocked
	resultA.Block(&resultB)

	// pull all users and set the flag to true next to any
	// user that was blocked by userA
	userBlockedQ := DB.
		Table("user_blocked_users").
		Select("1").
		Where("blocked_user_id = u.id")

	var users []User
	DB.Table("users u").
		Select("u.*, "+
			"(case when exists(?) then 'true' else 'false' end) as is_blocked",
			userBlockedQ,
		).
		Find(&users)

	if len(users) != 2 {
		t.Errorf("Invalid number of users pulled: %d; expected 2", len(users))
	}

	for _, u := range users {
		if u.ID == resultA.ID && u.IsBlocked == true {
			t.Errorf("User #%d should not be blocked by user #%d", resultA.ID, resultB.ID)
		}

		if u.ID == resultB.ID && u.IsBlocked != true {
			t.Errorf("User #%d should be blocked by user #%d", resultB.ID, resultA.ID)
		}
	}

	// ----------------------------------------------------------------------
	// `First` gives no error or complaints about missing column `is_blocked`
	var firstUser *User
	DB.Table("users").First(&firstUser)

	if firstUser.IsBlocked != false {
		t.Errorf("IsBlocked should be false as a default value")
	}

	chat.AddUser(&resultA)
	chat.AddUser(&resultB)

	var chatUsers []*User

	// ----------------------------------------------------------------------
	// When pulled as an association however, we get an error about missing
	// column `is_blocked`
	usersAssoc := DB.Model(chat).Association("Users")
	if err := usersAssoc.Find(&chatUsers); err != nil {
		t.Errorf("error pulling chat users: %v", err)
	}
}
