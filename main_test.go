package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: clickhouse

func TestGORM(t *testing.T) {

	// append data
	user := User{
		Name:     "ddddddddddddddddddddddddddddddddfsfdsfsd",
		Age:      0,
		Birthday: "dbdfbbbbbbbbbbbbbbbbbbbbbbbbbdfgfdgfdg",
		Account:  "apoiorpqwifpofpwiefpiweopfioeifoeifpwifpewifpwifpewifpewifpweifpiwefpiwepfiwpeifpwiefpiwe",
		Company:  "apoiorpqwifpofpwiefpiweopfioeifoeifpwifpewifpwifpewifpewifpweifpiwefpiwepfiwpeifpwiefpiweapoiorpqwifpofpwiefpiweopfioeifoeifpwifpewifpwifpewifpewifpweifpiwefpiwepfiwpeifpwiefpiweapoiorpqwifpofpwiefpiweopfioeifoeifpwifpewifpwifpewifpewifpweifpiwefpiwepfiwpeifpwiefpiweapoiorpqwifpofpwiefpiweopfioeifoeifpwifpewifpwifpewifpewifpweifpiwefpiwepfiwpeifpwiefpiweapoiorpqwifpofpwiefpiweopfioeifoeifpwifpewifpwifpewifpewifpweifpiwefpiwepfiwpeifpwiefpiweapoiorpqwifpofpwiefpiweopfioeifoeifpwifpewifpwifpewifpewifpweifpiwefpiwepfiwpeifpwiefpiweapoiorpqwifpofpwiefpiweopfioeifoeifpwifpewifpwifpewifpewifpweifpiwefpiwepfiwpeifpwiefpiweapoiorpqwifpofpwiefpiweopfioeifoeifpwifpewifpwifpewifpewifpweifpiwefpiwepfiwpeifpwiefpiwevvvapoiorpqwifpofpwiefpiweopfioeifoeifpwifpewifpwifpewifpewifpweifpiwefpiwepfiwpeifpwiefpiweapoiorpqwifpofpwiefpiweopfioeifoeifpwifpewifpwifpewifpewifpweifpiwefpiwepfiwpeifpwiefpiweapoiorpqwifpofpwiefpiweopfioeifoeifpwifpewifpwifpewifpewifpweifpiwefpiwepfiwpeifpwiefpiweapoiorpqwifpofpwiefpiweopfioeifoeifpwifpewifpwifpewifpewifpweifpiwefpiwepfiwpeifpwiefpiweapoiorpqwifpofpwiefpiweopfioeifoeifpwifpewifpwifpewifpewifpweifpiwefpiwepfiwpeifpwiefpiwe",
	}
	var users []User = make([]User, 0, 10000)
	for i := 0; i < 10000; i++ {
		users = append(users, user)
	}

	err := DB.CreateInBatches(users, len(users)).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	//execute will be fail; because exceed clickhouse query memory limit
	//but not return error
	var result []*User = make([]*User, 0)
	err = DB.Raw("select * from users limit 10000 settings max_memory_usage = 2;").Scan(&result).Error
	if err == nil {
		t.Errorf("error not capture")
	}
}
