package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	JsonField JsonField `gorm:"type:json;not null;" json:"jsonField"`
}

type JsonField struct {
	Array []string `json:"array"`
}

func (j *JsonField) Scan(src any) error {
	err := json.Unmarshal(src.([]byte), &j)
	if err != nil {
		return fmt.Errorf("json scan fail, err: %w", err)
	}

	for i, s := range j.Array {
		j.Array[i], _ = strings.CutSuffix(s, "1")
	}

	return nil
}

func (j JsonField) Value() (driver.Value, error) {
	clone := slices.Clone(j.Array)
	for i := range clone {
		clone[i] += "1" // 我想要修改（加密）原来的字段，I want to modify/encrypt the fields
	}

	data, err := json.Marshal(JsonField{
		Array: clone,
	})
	if err != nil {
		return nil, fmt.Errorf("json marshal fail, err: %w", err)
	}

	return data, nil
}

type SkipCustomMethod struct {
	Array []string `gorm:"type:json;not null;" json:"array"`
}

func (s *SkipCustomMethod) Scan(src any) error {
	return json.Unmarshal(src.([]byte), &s)
}

func (s SkipCustomMethod) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// ModelSkip to skip JsonField.Value and JsonField.Scan
type ModelSkip struct {
	JsonField SkipCustomMethod `json:"jsonField"`
}
