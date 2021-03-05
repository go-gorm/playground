package main

import (
	"database/sql"
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result = struct {
		RangeOne int
		RangeTwo int
	}{}

	// using these util function to get the values based at runtime based on current timestamp
	rangeOneStart, rangeOneEnd := getDailyRange(time.Now())
	rangeTwoStart, rangeTwoEnd := getYearlyRange(time.Now())

	err := DB.Model(user).
		Select("sum( if( created_at >= @rangeOneStart AND created_at < @rangeOneEnd, 1, 0 )) as range_one, "+
			"sum( if( created_at >= @rangeTwoStart AND created_at < @rangeTwoEnd, 1, 0 )) as range_two",
		sql.Named("rangeOneStart", rangeOneStart),
		sql.Named("rangeOneEnd", rangeOneEnd),
		sql.Named("rangeTwoStart", rangeTwoStart),
		sql.Named("rangeTwoEnd", rangeTwoEnd)).
		Where("created_at >= @startTime", sql.Named("startTime", rangeTwoStart)).
		Where("created_at< @endTime", sql.Named("endTime", rangeTwoEnd)).
		Scan(&result).Error

	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// Query is still valid when the named params are not replaced
	// hence it'll result the empty result
	// so check if the result is not 0 to assert this
	if result.RangeOne == 0 || result.RangeTwo == 0 {
		t.Errorf("Failed, query did not produce proper result")
	}
}

func getDailyRange(curTime time.Time) (string, string) {
	startTime := time.Date(
		curTime.Year(),
		curTime.Month(),
		curTime.Day(),
		0,
		0,
		0,
		0,
		curTime.Location())

	return startTime.Format("2006-01-02 15:04:05"), startTime.Add(24 * time.Hour).Format("2006-01-02 15:04:05")
}

func getYearlyRange(curTime time.Time) (string, string) {
	startTime := time.Date(
		curTime.Year(),
		1,
		1,
		0,
		0,
		0,
		0,
		curTime.Location())

	nextMonth := startTime.Add(365 * 24 * time.Hour)

	nextMonth = time.Date(
		nextMonth.Year(),
		1,
		1,
		0,
		0,
		0,
		0,
		nextMonth.Location())

	return startTime.Format("2006-01-02 15:04:05"), nextMonth.Format("2006-01-02 15:04:05")
}
