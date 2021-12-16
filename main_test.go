package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestGORMCompositeKey(t *testing.T) {
	// Insert New record
	prod := &CompositeKeyProduct{
		LanguageCode: 56,
		Code:         "56",
		Name:         "Test56",
	}

	nextProd := &SingleKeyProduct{
		LanguageCode: 200,
		Code:         "200",
		Name:         "Test200",
	}

	res := DB.Create(prod)
	if res.Error != nil {
		panic(res.Error)
	}

	res2 := DB.Create(nextProd)
	if res2.Error != nil {
		panic(res2.Error)
	}

	// fmt.Println(prod)
	assert.Equal(t, 56, prod.LanguageCode)
	assert.Equal(t, "56", prod.Code)
	assert.Equal(t, "Test56", prod.Name)
	assert.Equal(t, 1, prod.ProductPointID)

	// fmt.Println(nextProd)
	assert.Equal(t, 200, nextProd.LanguageCode)
	assert.Equal(t, "200", nextProd.Code)
	assert.Equal(t, "Test200", nextProd.Name)
	assert.Equal(t, 1, nextProd.ProductPointID)

}
