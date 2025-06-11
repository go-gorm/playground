package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{
		Name: "jinzhu",
		Pets: []*Pet{
			&Pet{
				Name: "puppy",
			},
			&Pet{
				Name: "kitty",
			},
			&Pet{
				Name: "puppy",
			},
		}}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Model(&User{}).
		Preload("Pets", func(db *gorm.DB) *gorm.DB {
			subscope := db.Select("MAX(id) as id").Model(&Pet{}).Group("name")
			return db.Where("id in (?)", subscope)
			// result ::
			// SELECT MAX(id) as id FROM `pets` WHERE `pets`.`user_id` = 1 AND id in (SELECT MAX(id) as id FROM `pets` WHERE `pets`.`user_id` = 1 AND id in () AND `pets`.`deleted_at` IS NULL GROUP BY `name`
			// expected ::
			// SELECT * FROM `pets` WHERE `pets`.`user_id` = 1 AND id in (SELECT MAX(id) as id FROM `pets` GROUP BY `name`)
		}).
		Where("id = ?", user.ID).
		Find(&result).Error; err != nil {
		t.Errorf("failed, get error: %v", err)
	}

}
