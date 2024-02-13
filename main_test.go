package main

import (
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestAssociationNotNullClear(t *testing.T) {
	type Profile struct {
		gorm.Model
		Number   string
		SSN      string `gorm:"uniqueIndex"`
		MemberID uint   `gorm:"not null"`
	}

	type Member struct {
		gorm.Model
		Profiles []Profile
	}

	DB.Migrator().DropTable(&Member{}, &Profile{})

	if err := DB.AutoMigrate(&Member{}, &Profile{}); err != nil {
		t.Fatalf("Failed to migrate, got error: %v", err)
	}

	member := &Member{
		Profiles: []Profile{{
			Number: "1",
		}, {
			Number: "2",
		}},
	}

	if err := DB.Create(&member).Error; err != nil {
		t.Fatalf("Failed to create test data, got error: %v", err)
	}

	if err := DB.Model(member).Association("Profiles").Clear(); err == nil {
		t.Fatalf("No error occurred during clearind not null association")
	}
}
