package main

import (
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	pst, _ := time.LoadLocation(("America/Los_Angeles"))
	est, _ := time.LoadLocation(("America/New_York"))
	now := time.Now()
	nowInPST := now.In(pst)
	nowInEST := now.In(est)

	/*
		user1 := User{Name: "jinzhu", Birthday: &now}
		user2 := User{Name: "jinzhu", Birthday: &nowInPST}
		user3 := User{Name: "jinzhu", Birthday: &nowInEST}
	*/
	x := DB.Explain("SELECT users WHERE birthday = ?", now)
	y := DB.Explain("SELECT users WHERE birthday = ?", nowInPST)
	z := DB.Explain("SELECT users WHERE birthday = ?", nowInEST)

	if x != y {
		t.Errorf("The SQL is not equal %s != %s", x, y)
	}

	if y != z {
		t.Errorf("The SQL is not equal %s != %s", y, z)
	}

	if x != z {
		t.Errorf("The SQL is not equal %s != %s", x, z)
	}
}
