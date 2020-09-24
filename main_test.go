package main

import (
	"log"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Test struct {
	ID    int
	Value string
}

func TestGORM(t *testing.T) {

	DB.Migrator().DropTable(Test{})
	DB.AutoMigrate(Test{})

	value := &Test{Value: "test"}

	if err := DB.Debug().Create(&value).Error; err != nil {
		t.Fatalf("Create: %s\n", err)
	}

	tests := []struct {
		name  string
		model interface{}
		isNil bool
	}{
		{name: "no model", isNil: true},
		{name: "ptr", model: &Test{}},
		{name: "not ptr", model: Test{}},
	}

	for _, test := range tests {
		scope := DB
		if !test.isNil {
			log.Println("model=", test.model)
			scope = scope.Model(test.model)
		}

		if err := scope.Save(&value).Error; err != nil {
			t.Errorf("%s: %s\n", test.name, err)
		}
	}
}
