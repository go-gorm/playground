package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	type Skill struct {
		gorm.Model
		ManagerID uint
		Name      string
	}

	type Manager struct {
		gorm.Model
		Name   string
		Skills []Skill
	}

	type Project struct {
		gorm.Model
		Name      string
		ManagerID uint
		Manager   Manager
	}

	if err := DB.Migrator().DropTable(
		&Skill{}, &Manager{}, &Project{},
	); err != nil {
		panic(err)
	}
	if err := DB.AutoMigrate(
		&Skill{}, &Manager{}, &Project{},
	); err != nil {
		t.Errorf("AutoMigrate failed: %v", err)
	}

	project := Project{
		Name: "foo project",
		Manager: Manager{
			Name: "foo manager",
			Skills: []Skill{
				{
					Name: "foo",
				},
				{
					Name: "bar",
				},
				{
					Name: "baz",
				},
			},
		},
	}
	DB.Create(&project)

	var result Project
	if err := DB.
		Joins("Manager").
		Preload("Manager.Skills").
		First(&result, project.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
