package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	// create 2 users, each with an one-2-many relationship
	user1 := User{Name: "Alice", Age: 34, Active: true,
		Pets: []*Pet{
			{Name: "Doggo"},
			{Name: "Pepe the frog"},
		}}
	user2 := User{Name: "Bob", Age: 35, Active: false, Pets: []*Pet{
		{Name: "Nyan cat"},
	}}

	if err := DB.Create(&user1).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if err := DB.Create(&user2).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// looking for a pet name 'Doggo'
	tx := DB.Where(Pet{Name: "Doggo"})

	//
	// SQL generated:
	// SELECT `pets`.`id`,`pets`.`created_at`,`pets`.`updated_at`,`pets`.`deleted_at`,`pets`.`user_id`,`pets`.`name` FROM `pets` JOIN users u on u.id = pets.user_id WHERE `pets`.`name` = "Doggo" AND u.name = "Alice" AND `pets`.`deleted_at` IS NULL ORDER BY `pets`.`id` LIMIT 1
	//
	// uncomment this and the test will pass:
	//
	// tx = tx.Joins("JOIN users u on u.id = pets.user_id").Where("u.name = ?", "Alice")
	//

	//
	// SQL generated:
	// SELECT `pets`.`id`,`pets`.`created_at`,`pets`.`updated_at`,`pets`.`deleted_at`,`pets`.`user_id`,`pets`.`name` FROM `pets` JOIN users u on u.id = pets.user_id WHERE `pets`.`name` = "Doggo" AND `pets`.`name` = "Alice" AND `pets`.`deleted_at` IS NULL ORDER BY `pets`.`id` LIMIT 1
	//
	// uncomment this and the test will fail:
	//
	tx = tx.Joins("JOIN users u on u.id = pets.user_id").Where(&User{Name: "Alice"})
	//

	pet := Pet{}
	if err := tx.First(&pet).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
