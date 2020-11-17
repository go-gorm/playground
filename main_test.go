package main

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	//	os.Setenv("GORM_DIALECT", "mysql")
	var Users [100]User
	for i := range Users {
		Users[i].Name = "Test"
		Users[i].Active = false
	}

	DB.Create(&Users)
	results := []User{}

	result := DB.Where("Active = ?", false).FindInBatches(&results, 10, func(tx *gorm.DB, batch int) error {

		for i := range results {

			results[i].Active = true

		}

		tx.Save(&results)
		fmt.Println(tx.RowsAffected) // number of records in this batch
		fmt.Println(batch)           // Batch 1, 2, 3

		return nil
	})
	fmt.Println(result.Error)        // returned error
	fmt.Println(result.RowsAffected) // processed records count in all batches

}
