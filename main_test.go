package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver
type JSON json.RawMessage

// gorm 查询
func (j *JSON) Scan(value interface{}) error {
	fmt.Println("ScanScanScanScan")
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	err := json.Unmarshal(bytes, j)
	return err
}

// gorm 入库
func (j JSON) Value() (driver.Value, error) {

	if len(j) == 0 {
		return nil, nil
	}
	fmt.Println(string(j))
	return json.RawMessage(j).MarshalJSON()
}

// json 自定义marshal
func (j JSON) MarshalJSON() ([]byte, error) {
	return json.RawMessage(j).MarshalJSON()
}

// json 自定义unmarshal
func (j *JSON) UnmarshalJSON(data []byte) error {
	*j = data
	return nil
}
type Jobs struct {
	ID                    uint
	RelabelConfigs       JSON           `json:"relabel_configs" gorm:"column:relabel_configs; type:json"`
}
func (Jobs) TableName() string {
	return "jobs"
}
type Targets struct {
	ID	    uint
	Target string `json:"target" gorm:"not null;unique"`
	JobID  uint   `json:"job_id" gorm:"column:job_id"`
	Job    Jobs   `json:"job"`
}

func TestGORM(t *testing.T) {
	job := Jobs{RelabelConfigs: []byte("{\"relabel_configs\": [{\"regex\": \"(.*):\\\\\\\\d+\", \"replacement\": \"1\", \"target_label\": \"host\", \"source_labels\": [\"__address__\"]}]}")}
	DB.AutoMigrate(Jobs{},Targets{})
	DB.Create(&job)
	var jobs []Jobs
        DB.Find(&jobs)
	for _, job := range jobs {
		fmt.Println("GetTarget前:RelabelConfigs", string(job.RelabelConfigs))
	}
	target := Targets{Target: "test1", JobID: 1}
	DB.Create(&target)
	var targets []Targets
	DB.Find(&targets)

	for _, job := range jobs {
		fmt.Println("GetTarget后:RelabelConfigs", string(job.RelabelConfigs))
	}
}
