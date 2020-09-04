package main

import "testing"

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Property struct {
	ID   int64  `gorm:"column:id;primaryKey"`
	Data string `gorm:"column:data"`
}

func TestGORM(t *testing.T) {
	DB.Migrator().AutoMigrate(&Property{})
	defer DB.Migrator().DropTable(&Property{})

	testData := []*Property{}
	for i := 1; i <= 100000; i++ {
		testData = append(testData, &Property{
			ID:   int64(i),
			Data: "aaa",
		})
	}

	if err := DB.Create(testData); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
