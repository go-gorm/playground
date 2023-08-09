package main

import (
	"gorm.io/gorm"
	"testing"
)

// Exact Example from https://gorm.io/docs/has_many.html#Declare

type MyUser struct {
	gorm.Model
	MyCreditCards []MyCreditCard
}

type MyCreditCard struct {
	gorm.Model
	Number   string
	MyUserID uint
}

func TestGORM(t *testing.T) {

	// This passes
	//if err := DB.AutoMigrate(&MyCreditCard{}, &MyUser{}); err != nil {
	//	t.Errorf("Failed, got error: %v", err)
	//}

	/* This fails. Error message:
	[0.489ms] [rows:0] ALTER TABLE `my_credit_cards` ADD CONSTRAINT `fk_my_users_my_credit_cards` FOREIGN KEY (`my_user_id`) REFERENCES `my_users`(`id`)
	--- FAIL: TestGORM (0.01s)
	    main_test.go:30: Failed, got error: Error 1824 (HY000): Failed to open the referenced table 'my_users'
	FAIL
	exit status 1
	FAIL    gorm.io/playground      0.364s
	*/
	if err := DB.AutoMigrate(&MyCreditCard{}, &MyUser{}); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
