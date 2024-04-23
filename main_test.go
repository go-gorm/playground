package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	users := []User{
		{
			Name: "Alexis",
			Company: &Company{
				Name: "Google",
				Addresses: []Addresses{
					{Name: "Mountain View"},
					{Name: "New York"},
				},
			},
		},
		{
			Name: "Jinzhu",
			Company: &Company{
				Name: "Google",
				Addresses: []Addresses{
					{Name: "Lyndhurst"},
					{Name: "California"},
				},
			},
		},
	}

	DB.Create(&users)

	var result []User
	customLogger.record = true
	err := DB.
		Joins("Company").
		Preload("Company.Addresses").
		Find(&result).Error

	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// queries should be like:
	// - SELECT `users`.`id`,`users`.`created_at`,`users`.`updated_at`,`users`.`deleted_at`,`users`.`name`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name`,`Company`.`user_id` AS `Company__user_id` FROM `users`
	//		LEFT JOIN `companies` `Company` ON `users`.`id` = `Company`.`user_id`
	//		WHERE `users`.`deleted_at` IS NULL
	//
	// - SELECT * FROM `companies` WHERE `companies`.`id` IN (1, 2)
	//
	// It should have 2 queries, not 3
	// Gorm is doing a n+1 query here, but it should not
	//
	// In version v1.25.6 the select of addresses where using a IN clause but have the problem of doing a second query with companies (https://github.com/go-gorm/gorm/issues/6715#issuecomment-1832455676)

	if len(queries) != 2 {
		t.Errorf("Failed, expect 2 queries, but got %v", len(queries))
	}
}
