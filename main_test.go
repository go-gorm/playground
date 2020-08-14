package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	initDb()

	users := []User{{Name: "jinzhu"}, {Name: "alice"}, {Name: "bob"}, {Name: "jinzhu"}, {Name: "charlie"}}
	errCreate := DB.Create(&users).Error
	if errCreate != nil {
		t.Errorf("Create of multiple users failed, got error: %v", errCreate)
		return
	}

	var result User
	results := make([]User, 0, 5)
	var count int64 = 0

	// Simply comment out the cases that you don't want to run

	// Case 0
	// Works, as Find does not need any ordering
	// SELECT * FROM "users" WHERE name = 'jinzhu'
	// SELECT count(1) FROM "users" WHERE name = 'jinzhu'
	errRetrieve := DB.Table("users").
		Where("name = 'jinzhu'").
		Find(&results).
		Count(&count).
		Error
	if errRetrieve != nil {
		t.Errorf("[0] Retrieving entries failed, got error: %v", errRetrieve)
		return
	}
	t.Logf("[0] count:   %d", count)
	t.Logf("[0] results: %v", results)

	// Case 1
	// Does not work, as first will insert a ORDER and postgres then needs to have GROUP BY specified
	// This was possible in v1
	// Comment this out to test the 2nd case
	// SELECT * FROM "users" WHERE name = 'jinzhu' ORDER BY "users"."id" LIMIT 1
	// SELECT count(1) FROM "users" WHERE name = 'jinzhu' ORDER BY "users"."id" LIMIT 1
	errRetrieve = DB.Table("users").
		Where("name = 'jinzhu'").
		First(&result).
		Count(&count).
		Error
	if errRetrieve != nil {
		t.Errorf("[1] Retrieving entries failed, got error: %v", errRetrieve)
		return
	}
	t.Logf("[1] count:  %d", count)
	t.Logf("[1] result: %v", result)

	// Case 2
	// Does mysteriously exit after the Count query (but the count query works), without even executing any defer statements!
	// SELECT count(1) FROM "users" WHERE name = 'jinzhu'
	// -
	defer func() {
		t.Logf("in defer")
	}()
	t.Logf("before 2nd case")
	errRetrieve = DB.Table("users").
		Where("name = 'jinzhu'").
		Count(&count).
		First(&result).
		Error
	if errRetrieve != nil {
		t.Errorf("[1] Retrieving entries failed, got error: %v", errRetrieve)
		return
	}
	t.Logf("after 2nd case")
	t.Logf("[2] count:  %d", count)
	t.Logf("[2] result: %v", result)

	// Case 3
	// Adding a GROUP BY statement will lead to the correct result, but is a bit cumbersome
	// SELECT * FROM "users" WHERE name = 'jinzhu' GROUP BY "id" ORDER BY "users"."id" LIMIT 1
	// SELECT count(1) FROM "users" WHERE name = 'jinzhu' GROUP BY "id" ORDER BY "users"."id" LIMIT 1
	errRetrieve = DB.Table("users").
		Where("name = 'jinzhu'").
		Group("id").
		First(&result).
		Count(&count).
		Error
	if errRetrieve != nil {
		t.Errorf("Retrieving entries failed, got error: %v", errRetrieve)
		return
	}
	t.Logf("[3] count:  %d", count)
	t.Logf("[3] result: %v", result)

}
