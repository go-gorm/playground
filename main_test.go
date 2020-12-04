package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

// XUserSpec is user-supplied "specification"
type XUserSpec struct {
	ID   uint64
	Name string

	Languages []string `gorm:"-"`
}

type XLanguage struct {
	Code string `gorm:"primarykey"`
}

// XUser is runtime object
type XUser struct {
	XUserSpec `godm:"embedded"`

	Languages []XLanguage `gorm:"many2many:XUserSpeak"`
}

func TestGORM(t *testing.T) {
	if err := DB.AutoMigrate(&XLanguage{}, &XUser{}); err != nil {
		//t.Fatal(err)
	}

	spec := XUserSpec{
		Name:      "jinzhu",
		Languages: []string{"English", "Chinese"},
	}

	user := XUser{XUserSpec: spec}

	for _, lang := range spec.Languages {
		user.Languages = append(user.Languages, XLanguage{Code: lang})
	}

	if err := DB.Create(&user).Error; err != nil {
		t.Fatal(err)
	}

	var result XUser
	if err := DB.Preload("Languages").First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if len(result.Languages) == 0 {
		t.Errorf("No languages assigned")
	}
}
