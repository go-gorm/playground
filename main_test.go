package main

import (
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Repo struct {
	ID     uint   `gorm:"primarykey"`
	URL    string `json:"RepoURL"`
	Status string `json:"RepoStatus"`
}

type DispatchRecord struct {
	ID                   uint `gorm:"primarykey"`
	PlaybookURL          string
	PlaybookDispatcherID string
}

type UpdateTransaction struct {
	ID              uint `gorm:"primarykey"`
	RepoID          *uint
	Repo            *Repo
	DispatchRecords []DispatchRecord `gorm:"many2many:updatetransaction_dispatchrecords"`
}

type DeviceUpdate struct {
	ID       uint `gorm:"primarykey"`
	Name     string
	UpdateID *uint
	Update   *UpdateTransaction
}

func TestGORM(t *testing.T) {
	_ = DB.AutoMigrate(&Repo{})
	_ = DB.AutoMigrate(&DispatchRecord{})
	_ = DB.AutoMigrate(&UpdateTransaction{})
	_ = DB.AutoMigrate(&DeviceUpdate{})

	deviceUpdate := DeviceUpdate{
		Name: "a devioce update",
		// Device: &Device{UUID: faker.UUIDHyphenated(), OrgID: orgID},
		Update: &UpdateTransaction{Repo: &Repo{URL: "http://some.example.com/content"}},
	}
	DB.Create(&deviceUpdate)

	var deviceUpdates []DeviceUpdate

	testCases := []struct {
		Name  string
		Query *gorm.DB
	}{
		{
			Name:  "2 preloads",
			Query: DB.Preload("Update.Repo").Preload("Update"),
		},
		{
			Name:  "2 Joins",
			Query: DB.Joins("Update.Repo").Joins("Update"),
		},
		{
			Name:  "3 Preloads",
			Query: DB.Preload("Update.Repo").Preload("Update").Preload("Update.DispatchRecords"),
		},
		// all bellow fails with panic with gorm version 1.25.7
		{
			Name:  "1 Preload 1 Join ",
			Query: DB.Preload("Update.Repo").Joins("Update"),
		},
		{
			Name:  "2 Preloads 1 Join ",
			Query: DB.Preload("Update.Repo").Joins("Update").Preload("Update.DispatchRecords"),
		},
		{
			Name:  "1 Preload 2 Joins ",
			Query: DB.Joins("Update.Repo").Joins("Update").Preload("Update.DispatchRecords"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if err := testCase.Query.Find(&deviceUpdates).Error; err != nil {
				t.Errorf("Failed, got error: %v", err)
			}
		})
	}
}
