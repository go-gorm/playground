package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite postgres

type Task struct {
	gorm.Model
	Content  string    `gorm:"not null; default:Empty Task" json:"content"`
	Deadline time.Time `json:"deadline,omitempty"`
	Priority uint      `json:"priority,omitempty"`
	Complete bool      `json:"complete,omitempty"`

	ParentTaskID *uint `gorm:"not null" json:"parent_task_id"`
	ParentTask   *Task `gorm:"foreignKey:ParentTaskID" json:"-"`
}

func TestGORM(t *testing.T) {
	require := require.New(t)
	require.NoError(DB.Migrator().DropTable(&Task{}))
	require.NoError(DB.Debug().AutoMigrate(&Task{}))

	var nonExistingTask uint = 30
	task := Task{
		Content:      "Test",
		Priority:     1,
		ParentTaskID: &nonExistingTask,
	}
	require.Error(DB.Debug().Create(&task).Error, "should return foreign key constraint error")
}
