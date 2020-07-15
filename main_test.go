package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if doc, count, err := result.GetPage(DB, 10, 1); err != nil {
		t.Errorf("Failed, got error: %v", err)
	} else {
		t.Log(doc, count)
	}
}

type UserPage struct {
	User
	CompanyName string `gorm:"-" json:"companyName"`
}

func (User) TableName() string {
	return "user"
}

func (Company) TableName() string {
	return "company"
}

func (e *User) GetPage(db *gorm.DB, pageSize int, pageIndex int) ([]UserPage, int64, error) {
	var doc []UserPage
	table := db.Select("user.*,company.name company_name").Table("user")
	table = table.Joins("left join company on user.company_id = company.id")

	var count int64

	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Count(&count)
	return doc, count, nil
}
