package main

import (
	"log"
	"testing"
	"time"
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
}

func TestUpdatedAt(t *testing.T) {
	DB.AutoMigrate(&BrokenUpdate{})

	model := BrokenUpdate{}
	err := DB.Create(&model).Error
	if err != nil {
		t.Errorf("Failed to create model: %v", err)

	}

	time.Sleep(time.Second * 1)

	// Save the created state
	before := model

	// We update a single field of the model
	log.Default().Print(model.UpdatedAt)
	tx := DB.Model(&model).Updates(BrokenUpdate{StringMember: "updated"})
	if tx.Error != nil {
		t.Errorf("Failed to update model: %v", tx.Error)
	}
	log.Default().Print(before.UpdatedAt)

	// We fetch the model again
	var retrieved BrokenUpdate

	result := DB.First(&retrieved, model.ID)
	if result.Error != nil {
		t.Errorf("Failed to find model: %v", result.Error)
	}
	log.Default().Print(retrieved.UpdatedAt)

	if result.RowsAffected == 0 {
		t.Errorf("Failed to find model")
	}

	// And then assert that the fields did/didn't change
	if !before.CreatedAt.Equal(retrieved.CreatedAt) {
		t.Errorf("CreatedAt should not change")
	}
	if !before.LastCheckinTime.Equal(retrieved.LastCheckinTime) {
		t.Errorf("LastCheckinTime should not change")
	}
	if before.SecretToken != retrieved.SecretToken {
		t.Errorf("SecretToken should not change")
	}
	if !before.DeletedAt.Time.Equal(retrieved.DeletedAt.Time) {
		t.Errorf("DeletedAt should not change")
	}
	if before.EmbeddedStruct != retrieved.EmbeddedStruct {
		t.Errorf("EmbeddedStruct should not change")
	}

	if before.ID != retrieved.ID {
		t.Errorf("ID should not change")
	}
	if before.Code != retrieved.Code {
		t.Errorf("Code should not change")
	}
	if before.JobName != retrieved.JobName {
		t.Errorf("JobName should not change")
	}
	// FIXME: UpdatedAt doesn't seem to actually change when .Updates() is used
	if before.UpdatedAt.Equal(retrieved.UpdatedAt) {
		t.Errorf("UpdatedAt should change")
	}
	if retrieved.StringMember != "updated" {
		t.Errorf("StringMember should change to 'updated'")
	}
}
