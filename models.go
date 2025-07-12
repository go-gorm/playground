package main

type Person struct {
	ID                         int64 `gorm:"primary_key"`
	FirstAndLastName           string
	CurrentEmployerID          *int64
	CurrentEmployerInformation *CompanyInformation `gorm:"foreignKey:CurrentEmployerID"`
}

type CompanyInformation struct {
	ID                         int64 `gorm:"primary_key"`
	Name                       string
	PreferredCompanyLanguageID *int64
	PreferredCompanyLanguage   *LanguageInformation `gorm:"foreignKey:PreferredCompanyLanguageID"`
}

type LanguageInformation struct {
	ID                int64 `gorm:"primarykey"`
	InternationalCode string
	Name              string
}
