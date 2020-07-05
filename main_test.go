package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Bugs struct {
	ID     int64
	Status BugStatus `gorm:"ForeignKey:Status;References:Status;Constraint:OnUpdate:CASCADE"`
}

type BugStatus struct {
	Status string `gorm:"PrimaryKey;type:varchar(22)"`
}

func TestGORM(t *testing.T) {
	if err := DB.AutoMigrate(new(BugStatus), new(Bugs)); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
