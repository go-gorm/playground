package main

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type MyUser struct {
	gorm.Model
	Name string `gorm:"type:varchar(20)"`
	Languages []MyLanguage `gorm:"many2many:user_languages;"`
}

type MyLanguage struct {
	gorm.Model
	Name string
}

func TestGORM(t *testing.T) {
	DB.AutoMigrate(&MyUser{})
	DB.AutoMigrate(&MyLanguage{})

	DB.Debug().Where("1 = 1").Delete(&MyUser{})
	DB.Debug().Where("1 = 1").Delete(&MyLanguage{})

	zh := MyLanguage{Name: "zh"}
	en := MyLanguage{Name: "en"}
	DB.Debug().Create(&zh)
	DB.Debug().Create(&en)

	var languages []MyLanguage
	DB.Find(&languages)
	fmt.Println(languages)

	user := MyUser{
		Name: "test_user",
		Languages: languages,
	}
	DB.Debug().Create(&user)

}
