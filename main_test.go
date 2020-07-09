package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Test1 struct {
	Map []string
}


func TestGORM(t *testing.T) {
	t.Run("Regular Map", func(t *testing.T) {
		err := DB.AutoMigrate(&Test1{})
		if err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
	})
}
