package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	//manager1 := User{Name: "manager1", Company: Company{Name: "company1", Country: "DE"}}
	//DB.Create(&manager1)
	//manager2 := User{Name: "manager2", Company: Company{Name: "company2", Country: "US"}}
	//DB.Create(&manager2)
	user1 := User{Name: "user1",
		Something: "1",
		Company:   Company{Name: "company1", Country: "DE"},
		Pets: []*Pet{
			{Name: "dog1"},
			{Name: "dog2"},
			{Name: "dog3"},
			{Name: "dog4"},
		}}
	DB.Create(&user1)
	user11 := User{Name: "user11",
		Something: "1",
		Company:   Company{Name: "company1", Country: "GB"},
		Pets: []*Pet{
			{Name: "dog1"},
			{Name: "dog22"},
			{Name: "dog33"},
			{Name: "dog4"},
		}}
	DB.Create(&user11)
	user2 := User{Name: "user2",
		Something: "2",
		Company:   Company{Name: "company2", Country: "US"},
		Pets: []*Pet{
			{Name: "dog4"},
			{Name: "dog6"},
			{Name: "dog7"},
			{Name: "dog8"},
		}}
	DB.Create(&user2)
	user22 := User{Name: "user22",
		Something: "2",
		Company:   Company{Name: "company2", Country: "DE"},
		Pets: []*Pet{
			{Name: "dog4"},
			{Name: "dog66"},
			{Name: "dog77"},
			{Name: "dog8"},
		}}
	DB.Create(&user22)

	query := DB.Table("users").
		Select("companies.country", "COUNT(DISTINCT users.name)").
		Joins("LEFT JOIN companies ON users.company_id = companies.id").
		Group("companies.country")
	query.Where(query.Where("users.name ILIKE ?", "%%").
		Or("companies.name ILIKE ?", "%%"))
	query.Where("users.something IN (?)", []string{"1"})

	type Result struct {
		Country string
		Count   int
	}

	var result []Result
	if err := query.Find(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	t.Logf("Result: %v", result)
	if len(result) != 2 {
		t.Errorf("Failed, unexpected length if results")
	}
	if result[0].Country != "DE" || result[0].Count != 1 {
		t.Errorf("Failed, not correct DE entry")
	}
	if result[1].Country != "GB" || result[1].Count != 1 {
		t.Errorf("Failed, not correct GB entry")
	}
}
