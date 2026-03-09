package main

import (
	"gorm.io/gorm"
	"gorm.io/playground/models"
)

func main() {
	db, err := OpenTestConnection()
	if err != nil {
		panic(err)
	}

	tx := db.Model(&models.User{}).Session(&gorm.Session{})

	tx.Where("id = ?", 111).Updates(&models.User{
		Model: gorm.Model{
			ID: 111,
		},
		Age: 666,
	})
	tx.Where("id = ?", 222).Updates(&models.User{ // Query will be polluted with `id` = 111
		Model: gorm.Model{
			ID: 222,
		},
		Age: 777,
	})
	tx.Where("id = ?", 333).Updates(&models.User{ // Query will now be polluted with `id` = 222
		Model: gorm.Model{
			ID: 333,
		},
		Age: 888,
	})
}
