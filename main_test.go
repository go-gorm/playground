package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	// Migrate the schema
	_ = DB.AutoMigrate(&Language{})
	l1 := &Language{Code: "Chinese", Name: "可能是中文"}
	l2 := &Language{Code: "English", Name: "foo"}
	DB.Create(l1)
	DB.Create(l2)

	var languages1 []Language
	queryConds := []string{"Chinese"}
	err := DB.Find(&languages1, queryConds).Error // find language with code Chinese
	if err != nil {
		t.Errorf("err is %v, expect : nil", err)
	}
	if len(languages1) != 1 {
		t.Errorf("expect found one res but get :%v", len(languages1))
	} else {
		assert.Equal(t, l1.Code, languages1[0].Code)
		assert.Equal(t, l1.Name, languages1[0].Name)
	}

	var languages2 []Language
	queryConds = []string{}
	err = DB.Where("code IN ?", queryConds).Find(&languages2).Error // where find language with emptyConds
	if err != nil {
		t.Errorf("err is %v, expect : nil", err)
	}
	if len(languages2) != 0 {
		t.Errorf("expect found zeor res but get :%v", len(languages2))
	}


	var languages3 []Language
	queryConds = []string{}
	err = DB.Find(&languages3, queryConds).Error // find languages with empetyConds
	if err != nil {
		t.Errorf("err is %v, expect : nil", err)
	}
	if len(languages3) != 0 {
		t.Errorf("expect found zeor res but get :%v", len(languages3))
	}
}
