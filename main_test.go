package main

import (
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver
type IncomeStatistic struct {
	Date   time.Time `gorm:"column:date;primaryKey;type:date"`
	Amount int
	Count  int
}

func (i *IncomeStatistic) TableName() string {
	return "income_statistics"
}

type ExpenseStatistic struct {
	Date   time.Time `gorm:"column:date;primaryKey;type:date"`
	Amount int
	Count  int
}

func (e *ExpenseStatistic) TableName() string {
	return "expense_statistics"
}
func TestGORM(t *testing.T) {
	err := DB.AutoMigrate(&IncomeStatistic{}, &ExpenseStatistic{})
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	// 只有income_statistics表有数据
	DB.Create(&IncomeStatistic{Date: time.Now(), Count: 3, Amount: 3})
	query := DB.Where(
		"date = ?", time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local),
	)
	var result IeStatistic
	err = query.Select("SUM(amount) AS Amount,SUM(count) AS Count").Model(&IncomeStatistic{}).Scan(&result.Income).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	err = query.Select("SUM(amount) AS Amount,SUM(count) AS Count").Model(&ExpenseStatistic{}).Scan(&result.Expense).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	// 查询正确结果
	var expectedResult IeStatistic
	err = DB.Where(
		"date = ?", time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local),
	).Select("SUM(amount) AS Amount,SUM(count) AS Count").Table((&IncomeStatistic{}).TableName()).Scan(&expectedResult.Income).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	err = DB.Where(
		"date = ?", time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local),
	).Select("SUM(amount) AS Amount,SUM(count) AS Count").Table((&ExpenseStatistic{}).TableName()).Scan(&expectedResult.Expense).Error
	if result != expectedResult {
		t.Error("预期结果不符!", "预期结果：", expectedResult, "运行结果", result)
	}
}

type AmountCount struct {
	Amount int
	Count  int
}

type IeStatistic struct {
	Income  AmountCount
	Expense AmountCount
}
