package main

import (
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
type Targets struct {
	Model
	Target string `json:"target" gorm:"not null;unique"`
	JobID  uint   `json:"job_id" gorm:"column:job_id"`
	Job    Jobs   `json:"job"`
}

func TestGORM(t *testing.T) {
	job := Jobs{RelabelConfigs: []byte("{\"relabel_configs\": [{\"regex\": \"(.*):\\\\\\\\d+\", \"replacement\": \"1\", \"target_label\": \"host\", \"source_labels\": [\"__address__\"]}]}")}
	
	DB.Create(&job)
	var jobs []model.Jobs
        DB.Find(&jobs)
	for _, job := range jobs {
		fmt.Println("GetTarget前:RelabelConfigs", string(job.RelabelConfigs))
	}
	target := Targets{Target: "test1", JobID: 1}
	DB.Create(&target)
	var targets = []Targets
	DB.Find(&target)
	
	for _, job := range jobs {
		fmt.Println("GetTarget后:RelabelConfigs", string(job.RelabelConfigs))
	}
}
