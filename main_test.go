package main

import (
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

	familiar := Familiar{Name: "fam"}
	familiarIDBefore := familiar.Id
	DB.Create(&familiar)
	familiarIDAfter := familiar.Id
	if familiarIDBefore == familiarIDAfter {
		t.Errorf("Familiar ID should be created")
	}

	meeting := Meeting{IsLive: true}
	DB.Create(&meeting)
	familiars := []Familiar{familiar}
	meetingFamiliars := meeting.Familiars
	if len(meetingFamiliars) > 0 {
		t.Errorf("Meeting should not have any familiars yet")
	}

	DB.Model(&meeting).Association("Familiars").Append(familiars)
	meetingFamiliars = meeting.Familiars
	if len(meetingFamiliars) != 1 {
		t.Errorf("Meeting should have a familiar now")
	}

	familiarInMeetingID := meetingFamiliars[0].Id

	familiarIDAfterMeetingAdd := familiar.Id
	if familiarIDAfterMeetingAdd != familiarIDAfter {
		t.Errorf("Familiar ID should not be changed")
	}

	if familiarInMeetingID != familiarIDAfter {
		t.Errorf("Familiar ID should be the same. Meeting familiar ID: %v, familiar ID: %v", familiarInMeetingID, familiarIDAfter)
	}
}
