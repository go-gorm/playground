package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	users := []User{
		{Name: "subquery_having_1", Age: 10},
		{Name: "subquery_having_2", Age: 20},
		{Name: "subquery_having_3", Age: 30},
		{Name: "subquery_having_4", Age: 40},
	}
	DB.Create(&users)

	var results []User
	//OK
	DB.Select("AVG(age) as age, name").Table("users").Where("name LIKE ?", "subquery_having%").Group("name").Having("AVG(age) > (?)", DB.
		Select("AVG(age)").Where("name LIKE ?", "subquery_having%").Table("users")).Find(&results)
	if len(results) != 2 {
		t.Errorf("Two user group should be found, instead found %d", len(results))
	}

	results = nil
	//Failed
	DB.Select("AVG(age) as age, name").Table("users").Where("name LIKE ?", "subquery_having%").Group("users.name").Having("AVG(age) > (?)", DB.
		Select("AVG(age)").Where("name LIKE ?", "subquery_having%").Table("users")).Find(&results)

	if len(results) != 2 {
		t.Errorf("Two user group should be found, instead found %d", len(results))
	}
}


