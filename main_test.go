package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var beforeHookString = "before hook"

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Home struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey"`
	Address   *string   `json:"address"`
	Zip       *int      `json:"zip"`
	PriceArea *string   `json:"price_area"`
	Type      *string   `json:"type"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type HomeUpdate struct {
	Zip        *int    `json:"zip"`
	Address    *string `json:"address"`
	PriceArea  *string `json:"price_area"`
	Type       *string `json:"type"`
	created_at time.Time
	updated_at time.Time
}

func (home *Home) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("price_area", &beforeHookString)
	return nil
}

func TestGORM(t *testing.T) {
	var id = uuid.New()

	DB.AutoMigrate(&Home{})

	home := Home{
		Id: id,
	}

	dbErr := DB.Create(&home).Error
	if dbErr != nil {
		fmt.Println(dbErr.Error())
		return
	}

	tmpString := "test"
	updatePayload := HomeUpdate{
		Address: &tmpString,
	}

	dbErr = DB.Model(&home).Updates(&updatePayload).Error
	if dbErr != nil {
		fmt.Println(dbErr.Error())
		return
	}

	updatedHome := Home{
		Id: id,
	}

	dbErr = DB.First(&updatedHome).Error
	if dbErr != nil {
		fmt.Println(dbErr.Error())
		return
	}

	fmt.Printf("%+v\n", updatePayload)

	if *updatePayload.Type == beforeHookString {
		t.Errorf("Failed, wrong column is updated in before hook")
		return
	}

	if updatedHome.PriceArea == nil {
		t.Errorf("Failed, column is not updated in hook")
		return
	}

	if *updatedHome.PriceArea != beforeHookString {
		t.Errorf("Failed, column in hook not updated correctly")
		return
	}
}
