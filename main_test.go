package main

import (
	"testing"
	//"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{
		Name: "jinzhu",
		Languages: []Language{
			Language{
				Code: "EN",
				Name: "English",
			},
		},
	}

	DB.Set("user", "gabriel").Create(&user)

	// I also tried Scopes() but didn't work too
	// DB.Scopes(func(db *gorm.DB) *gorm.DB {
	//     return db.Set("user", "gabriel")
	// }).Create(&user)

	var result User
	if err := DB.Preload("Languages").First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
