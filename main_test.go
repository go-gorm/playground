package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

//TODO: Run migrations
//Create a view
//Re-Run Migrations and watch it fail boooooyaaaah
func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)
	err := DB.Exec("create view test_view_1 as select * from users").Error
	if err != nil {
		t.Errorf("Failed to create a test view")
	}
	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	err = DB.AutoMigrate(&User{})
	if err != nil {
		t.Errorf("Migration fails")
	}
	err = DB.Exec("drop view test_view_1").Error
	if err != nil {
		t.Errorf("Failed to drop view")
	}
}
