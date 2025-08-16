package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	DB.AutoMigrate(&Aster{}, &Bold{})
	//DB.Delete(&Bold{}, DB.Where("true = ?", true))
	//DB.Delete(&Aster{}, DB.Where("true = ?", true))

	dataTest := Aster{
		Name: "Sergio",
		Bolds: &[]Bold{
			{Tech: "Go"},
			{Tech: "Python"},
		},
	}

	DB.Create(&dataTest)

	var result *Aster
	if err := DB.Joins("Bolds").First(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
