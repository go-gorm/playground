package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	// Add data
	//startTime := time.Now()
	//// data 1
	//c1 := Conversation{
	//	ConversationId: "test-1",
	//	StartTime:      &startTime,
	//	EndTime:        nil,
	//	Title:          "1",
	//	Status:         "1",
	//	CreatedBy:      "user-1",
	//	MessageCount:   0,
	//	Language:       "",
	//	TokensConsumed: 0,
	//}
	//// data 2
	//c2 := Conversation{
	//	ConversationId: "test-2",
	//	StartTime:      &startTime,
	//	EndTime:        nil,
	//	Title:          "1",
	//	Status:         "1",
	//	CreatedBy:      "user-1",
	//	MessageCount:   0,
	//	Language:       "",
	//	TokensConsumed: 0,
	//}
	//DB.Create(&c1)
	//DB.Create(&c2)

	var cs []*Conversation
	err := DB.Where("created_by = ?", "user-1").Find(&cs).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	t.Log(cs[0].ConversationId)

}
