package main

import (
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	type TimeStamps struct {
		CreatedAt time.Time `json:"created_at,omitempty"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
	}

	type OrderBase struct {
		TimeStamps
		ID uint `json:"id,omitempty" gorm:"primarykey"`
	}

	type Order struct {
		OrderBase
		Reference string `json:"reference,omitempty"`
	}

	DB.Migrator().DropTable(&Order{})
	DB.AutoMigrate(&Order{})

	newOrder := Order{
		OrderBase: OrderBase{},
		Reference: "111",
	}
	DB.Save(&newOrder)

	if newOrder.ID == 0 {
		t.Errorf("order id should not be 0")
	}
}
