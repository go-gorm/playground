package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Result struct {
	Total           int64
	TodayTotal      int64
	TodayTotalMoney int64
}

func TestGORM(t *testing.T) {
	// timestamp
	var start, stop int64 = 1615219200, 1615305599

	var result = new(Result)

	err := DB.Table("order").Select([]string{
		"COUNT(id) AS total",
		"SUM(if(create_time BETWEEN ? AND ?, 1, 0)) as today_total",
		"SUM(if(create_time BETWEEN ? AND ?, money, 0)) as today_total_money",
	}, start, stop, start, stop).Scan(result).Error

	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	err = DB.Table("order").Select([]string{
		"COUNT(id) AS total",
		"SUM(if(create_time BETWEEN @start AND @stop, 1, 0)) as today_total",
		"SUM(if(create_time BETWEEN @start AND @stop, money, 0)) as today_total_money",
	}, start, stop).Scan(result).Error

	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
