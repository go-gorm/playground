package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	iv := DB.
		Table(`table_invoices`).
		Select(`seller, SUM(total) as total, SUM(paid) as paid, SUM(balance) as balance`).
		Group(`seller`)

	tx := DB.
		Table(`table_employees`).
		Select(`id, name, iv.total, iv.paid, iv.balance`).
		Joins(`LEFT JOIN (?) AS iv ON iv.seller = table_employees.id`, iv).
		Scan(&user)
	fmt.Println(tx.Error)
}
