package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql

type Settings struct {
	ID uint `gorm:"type:mediumint;unsigned;primary_key;AUTO_INCREMENT;unsigned" json:"ID"`
	DiscountCode string `gorm:"type:varchar(255);not null" json:"discount_code"`
}



func TestGORM(t *testing.T) {
	DB.Debug().AutoMigrate(&Settings{})

	settings := Settings{
		DiscountCode: "test",
	}
	
	DB.Create(&settings)
	
	find := Settings{
		DiscountCode: "test",
	}
	err := DB.Debug().Where(&find).First(&find).Error
	if err != nil {
		t.Errorf("Failed, got error: %s", err.Error())
		return
	}
	if find.ID == 0 {
		t.Errorf("Failed, did not get correct struct %v", find)
	}
}
