package main

import (
	"testing"

	"gorm.io/gorm"
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
}

func TestIncorrectRowsAffected(t *testing.T) {
	user := User{Name: "jinzhu2", Age: 25, Active: true}
	DB.Create(&user)
	user = User{Name: "jinzhu3", Age: 25, Active: true}
	DB.Create(&user)
	user = User{Name: "jinzhu4", Age: 25, Active: true}
	DB.Create(&user)

	condition := gorm.Expr("users.age = ? AND active = ?", 25, true)
	res := DB.Exec(`UPDATE users SET birthday = '2025-03-07' WHERE ?`, condition)
	t.Logf("res.RowsAffected - Expected = 3, Actual = %d", res.RowsAffected)
	if res.RowsAffected != 3 {
		t.Fail()
	}
	condition2 := gorm.Expr("users.age = ? AND active = ? AND company_id IS NULL", 25, true)
	res2 := DB.Exec(`UPDATE users SET birthday = '2025-03-07' WHERE ?`, condition2)
	t.Logf("res.RowsAffected - Expected = 3, Actual = %d", res2.RowsAffected)
	if res2.RowsAffected != 3 {
		t.Fail()
	}
	DB.Exec(`DELETE FROM users WHERE age = 25`)
}

func TestIncorrectRowsAffectedReversed(t *testing.T) {
	user := User{Name: "jinzhu2", Age: 25, Active: true}
	DB.Create(&user)
	user = User{Name: "jinzhu3", Age: 25, Active: true}
	DB.Create(&user)
	user = User{Name: "jinzhu4", Age: 25, Active: true}
	DB.Create(&user)

	condition2 := gorm.Expr("users.age = ? AND active = ? AND company_id IS NULL", 25, true)
	res2 := DB.Exec(`UPDATE users SET birthday = '2025-03-07' WHERE ?`, condition2)
	t.Logf("res.RowsAffected - Expected = 3, Actual = %d", res2.RowsAffected)
	if res2.RowsAffected != 3 {
		t.Fail()
	}
	condition := gorm.Expr("users.age = ? AND active = ?", 25, true)
	res := DB.Exec(`UPDATE users SET birthday = '2025-03-07' WHERE ?`, condition)
	t.Logf("res.RowsAffected - Expected = 3, Actual = %d", res.RowsAffected)
	if res.RowsAffected != 3 {
		t.Fail()
	}
	DB.Exec(`DELETE FROM users WHERE age = 25`)
}
