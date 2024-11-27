package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User struct {
	ID          int64           `json:"id"`
	Description UserDescription `json:"description"`
}

type UserDescription json.RawMessage

func (j *UserDescription) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("Failed to unmarshal JSONB value: %+v", value)
	}

	result := json.RawMessage{}
	err := json.Unmarshal([]byte(str), &result)
	*j = UserDescription(result)

	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j UserDescription) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	str := string(j)
	return str, nil
}
