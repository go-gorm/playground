package main

import (
	"encoding/json"
	"strings"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	jsonData := `{"name": "json-1", "attributes": ["tag1", "tag2"]}`
	var jsonMap map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(jsonData)).Decode(&jsonMap); err != nil {
		t.Fatal(err)
	}

	user := &UserWithJSON{}
	DB.Where(&UserWithJSON{Name: "json-1"}).Assign(jsonMap).FirstOrCreate(user)

	var result UserWithJSON
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
