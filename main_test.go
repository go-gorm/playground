package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Project struct {
	gorm.Model
	Name      string
	Workflows []*Workflow
}

type Workflow struct {
	gorm.Model
	Name      string
	ProjectID uint `gorm:"index"`
	Tasks     []*Task
}

type Task struct {
	gorm.Model
	Name       string
	WorkflowID uint    `gorm:"index"`
	PrevDeps   []*Task `gorm:"many2many:task_deps;joinForeignKey:NextID;joinReferences:PrevID"`
	NextDeps   []*Task `gorm:"many2many:task_deps;joinForeignKey:PrevID;joinReferences:NextID"`
}

func init() {
	_ = DB.AutoMigrate(&Project{}, &Workflow{}, &Task{})
}

func prepare() (pid, wid, tid, nid uint, err error) {
	// 1. create a project
	createdProject := &Project{Name: "Project_A"}
	DB.Create(&createdProject)
	pid = createdProject.ID

	// 2. create a workflow belonging to a project
	createdWorkflow := &Workflow{Name: "Workflow_A", ProjectID: createdProject.ID}
	DB.Create(&createdWorkflow)
	wid = createdWorkflow.ID

	// 3. create a task belonging to a workflow
	createdTask := &Task{Name: "Task_A", WorkflowID: createdWorkflow.ID}
	DB.Create(&createdTask)
	tid = createdTask.ID

	// 4. create a task with prev deps
	createdTaskWithPrevDeps := &Task{Name: "Task_B", WorkflowID: createdWorkflow.ID}
	createdTaskWithPrevDeps.PrevDeps = append(createdTaskWithPrevDeps.PrevDeps, &Task{
		Model: gorm.Model{ID: createdTask.ID},
	})
	DB.Create(&createdTaskWithPrevDeps)
	nid = createdTaskWithPrevDeps.ID

	if pid == 0 || wid == 0 || tid == 0 || nid == 0 {
		return 0, 0, 0, 0, fmt.Errorf("failed to prepare data")
	}

	return pid, wid, tid, nid, nil
}

func TestDeleteProjectWithWorkflowTask(t *testing.T) {
	pid, wid, tid, nid, err := prepare()
	_ = assert.Nil(t, err)

	var p Project
	err = DB.First(&p, pid).Error
	_ = assert.Nil(t, err)

	// Check Point: if Project_A has been deleted, Workflow_A and Task_A should be deleted too
	err = DB.Select("Workflows", "Workflows.Tasks", "Workflows.Tasks.NextDeps").Delete(&p).Error
	_ = assert.Nil(t, err)

	err = DB.First(&Project{}, pid).Error
	_ = assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = DB.First(&Workflow{}, wid).Error
	_ = assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = DB.First(&Task{}, tid).Error
	_ = assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = DB.First(&Task{}, nid).Error
	_ = assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestDeleteWorkflowWithTask(t *testing.T) {
	_, wid, tid, nid, err := prepare()
	_ = assert.Nil(t, err)

	var w Workflow
	err = DB.First(&w, wid).Error
	_ = assert.Nil(t, err)

	// Check Point: if Workflow_A has been deleted, Task_A and Task_B should be deleted too
	err = DB.Select("Tasks", "Tasks.NextDeps").Delete(&w).Error
	_ = assert.Nil(t, err)

	err = DB.First(&Workflow{}, wid).Error
	_ = assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = DB.First(&Task{}, tid).Error
	_ = assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = DB.First(&Task{}, nid).Error
	_ = assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestDeleteTasksWithNextDeps(t *testing.T) {
	_, _, tid, nid, err := prepare()
	_ = assert.Nil(t, err)

	var task Task
	err = DB.First(&task, tid).Error
	_ = assert.Nil(t, err)

	// Check Point: if Task_A has been deleted, Task_B should be deleted too
	err = DB.Select("NextDeps").Delete(&task).Error
	_ = assert.Nil(t, err)

	err = DB.First(&Task{}, tid).Error
	_ = assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = DB.First(&Task{}, nid).Error
	_ = assert.Equal(t, gorm.ErrRecordNotFound, err)
}
