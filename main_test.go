package main

import (
	"database/sql"
	"log"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type Result struct {
	User
	UserLanguage
	Language
	Pet
}

func TestGORM(t *testing.T) {
	if err := DB.Create(&User{Name: "jinzhu"}).Error; err != nil {
		panic(err)
	}
	if err := DB.Create(&Language{Name: "ZH"}).Error; err != nil {
		panic(err)
	}
	if err := DB.Create(&UserLanguage{UserID: 1, LanguageID: 1, Skilled: sql.NullBool{Bool: true, Valid: true}}).Error; err != nil {
		panic(err)
	}
	if err := DB.Create(&Pet{UserID: 1, Name: "mimi"}).Error; err != nil {
		panic(err)
	}
	results := query()
	for _, result := range results {
		if result.Pet.UserID == 0 {
			panic("result.Pet.UserID == 0")
		} else if result.Pet.Name != "mimi" {
			log.Fatalf("result.Pet.Name != \"mimi\": %s", result.Pet.Name)
		}
	}
	if err := DB.Create(&Pet{UserID: 1, Name: "wang"}).Error; err != nil {
		panic(err)
	}
	results = query()
	for _, result := range results {
		if result.Language.Name != "ZH" {
			log.Fatalf("result.Language.Name != \"ZH\": %s", result.Language.Name)
		}
		if result.Pet.UserID == 0 {
			panic("result.Pet.UserID == 0")
		}
	}
}

func query() []Result {
	results := make([]Result, 0, 8)
	if err := DB.Select("user.*, user_language.*, language.*, pet.*").Table("user").
		Joins("JOIN user_language ON user_language.user_id = user.id").
		Joins("JOIN language ON language.id = user_language.language_id").
		Joins("LEFT OUTER JOIN pet ON pet.user_id = user.id").Find(&results).Error; err != nil {
		panic(err)
	}
	return results
}
