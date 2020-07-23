package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver
type Person struct {
	Id   int64 `gorm:"type:bigint(20);NOT NULL;PRIMARY_KEY;"`
	Name string
}
func TestGORM(t *testing.T) {
	person := Person{Id: 1231231231231, Name: "test"}
	DB.AutoMigrate(&Person{})
	if err := DB.FirstOrCreate(&person, person).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
		return
	}
	if person.Id == 0 {
		t.Errorf("Failed, got error: id=0")
		return
	}

}
