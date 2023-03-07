package main

import (
	"testing"
	"time"

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

	DB.AutoMigrate(&EvalJob{})

	DB.Create(&EvalJob{
		Model: gorm.Model{
			ID:        0,
			CreatedAt: time.Now().Add(-24 * time.Hour),
			UpdatedAt: time.Now().Add(-24 * time.Hour),
		},
		PSM:      "inf.lidar.api",
		TaskID:   "taskid3",
		Operator: "xiehengjian",
		JobName:  "jobname",
		// Status:      models.EvalFailed,
		// RelatedData: map[RelatedDataType]interface{}{},
		// ProfitTimes: make(map[string]*ProfitTime),
	})
}

type EvalJob struct {
	gorm.Model
	PSM      string
	TaskID   string // 任务ID
	Operator string // 该job的发起人
	JobName  string
	// RespResult  ProfitDetail                    // 收益结果
	// Status      EvalJobStatus                   // 状态，1 已提交，2 计算中，3 暂停，4 失败 5 成功
	RelatedData map[RelatedDataType]interface{} `gorm:"serializer:json;type:text"`
	Region      string
	Internal    int // 1代表内部计算，0代表外部计算
	// 新增字段
	ProfitTimes map[string]*ProfitTime `gorm:"serializer:json"` // vregion和对应的计算时间
	ReportData  Report                 `gorm:"serializer:json"` // 收益报告详细信息
}

type RelatedDataType int

type ProfitTime struct{}

type Report struct{}
