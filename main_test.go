package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres

func TestGORM(t *testing.T) {
	languageEnglish := LanguageInformation{
		InternationalCode: "en-us",
		Name:              "English",
	}
	DB.Create(&languageEnglish)

	company1 := CompanyInformation{
		Name:                       "Company 1",
		PreferredCompanyLanguageID: &languageEnglish.ID,
	}
	DB.Create(&company1)

	ceoPerson := Person{
		FirstAndLastName:  "John Smith",
		CurrentEmployerID: &company1.ID,
	}
	DB.Create(&ceoPerson)

	var result Person
	err := DB.Model(&Person{}).
		Joins("CurrentEmployerInformation").
		Joins("CurrentEmployerInformation.PreferredCompanyLanguage").
		Limit(1).
		First(&result).
		Error

	if err != nil {
		t.Fatalf("Got DB error %v", err)
	}

	// All these should pass
	if result.FirstAndLastName != "John Smith" {
		t.Fatal("FirstAndLastName should be John Smith")
	}
	if result.CurrentEmployerInformation == nil {
		t.Fatal("CurrentEmployerInformation should not be nil")
	}

	if result.CurrentEmployerInformation.Name != "Company 1" {
		t.Fatal("CurrentEmployerInformation.Name should be Company 1")
	}

	if result.CurrentEmployerInformation.PreferredCompanyLanguage == nil {
		t.Fatal("CurrentEmployerInformation.PreferredCompanyLanguage should not be nil")
	}
	if result.CurrentEmployerInformation.PreferredCompanyLanguage.Name != "English" {
		t.Fatal("CurrentEmployerInformation.PreferredCompanyLanguage should be English")
	}

	/*
		This is the one that is demonstrating the specific bug with Postgres:
		The generated SQL query for this column is:
		`CurrentEmployerInformation__PreferredCompanyLanguage`.`international_code` AS `CurrentEmployerInformation__PreferredCompanyLanguage__international_code`,

		CurrentEmployerInformation__PreferredCompanyLanguage__international_code is larger than 63 characters long. Postgres
		accepts this in the query correctly, but returns a truncated identifier for this column that is 63 characters long.

		Gorm does not map the truncated column identifier to the intended column, leaving it empty.
		Note that this happens silently; other columns are still populated.

		Using the `NamingStrategy` with max length = 63 does not affect the identifier length in query building, and
		thus does not address this issue.
	*/
	if result.CurrentEmployerInformation.PreferredCompanyLanguage.InternationalCode != "en-us" {
		t.Fatalf("CurrentEmployerInformation.PreferredCompanyLanguage should be en-us, but was `%v`", result.CurrentEmployerInformation.PreferredCompanyLanguage.InternationalCode)
	}
}
