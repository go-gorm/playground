package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
    l1 := Language{Code:"English",Name:"English"}
    l2 := Language{Code:"French",Name:"French"}
    user1 := User{Name: "jinzhu", Age: 1, Languages:[]Language{l1, l2}}
    user2 := User{Name: "jinzhu", Age: 2, Languages:[]Language{l1, l2}}

	DB.Create(&l1)
	DB.Create(&l2)
	DB.Create(&user1)
	DB.Create(&user2)

    count := int64(0)

	var result []User
	if err := DB.Model(&User{}).
        Joins("JOIN user_speaks ON user_speaks.user_id = users.id").
        Joins("JOIN languages ON user_speaks.language_code = languages.code").
        Distinct("users.name, users.age").
        Count(&count).
        Find(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

    if len(result) != 2 {
		t.Errorf("Failed, expected only two users, got %d", len(result))
    }

    if count != 2 {
		t.Errorf("Failed, expected only two users in count, got %d", count)
    }
}
