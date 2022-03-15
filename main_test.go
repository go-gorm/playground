package main

import (
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	var r interface{}
	DB.Raw("DELETE FROM article_names").Scan(r)
	DB.Raw("DELETE FROM articles").Scan(r)

	articles := []*Article{
		{
			Code: "123",
			ArticleNames: []*ArticleName{
				{Name: "Jacket", Language: "en"},
				{Name: "Giacca", Language: "it"},
			},
		},
		{
			Code: "456",
			ArticleNames: []*ArticleName{
				{Name: "Shirt", Language: "en"},
				{Name: "Maglietta", Language: "it"},
			},
		},
	}

	err := DB.Clauses(clause.OnConflict{UpdateAll: true}).
		Session(&gorm.Session{CreateBatchSize: 1000, FullSaveAssociations: true}).
		Create(&articles).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	err = DB.Preload("ArticleNames").Find(&articles).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if len(articles) != 2 {
		t.Errorf("Expected 2 articles, found %d", len(articles))
	}
	if articles[0].Code != "123" {
		t.Errorf("Expected article 1 with code 123, found %s", articles[0].Code)
	}
	if len(articles[0].ArticleNames) != 2 {
		t.Errorf("Expected article 1 with 2 names, found %d", len(articles[0].ArticleNames))
	}
	if articles[0].ArticleNames[0].Name != "Jacket" {
		t.Errorf("Expected Jacket, get %s", articles[0].ArticleNames[0].Name)
	}
	if articles[0].ArticleNames[1].Name != "Giacca" {
		t.Errorf("Expected Giacca, get %s", articles[0].ArticleNames[1].Name)
	}
	if articles[1].Code != "456" {
		t.Errorf("Expected article 2 with code 456, found %s", articles[1].Code)
	}
	if len(articles[1].ArticleNames) != 2 {
		t.Errorf("Expected article 2 with 2 names, found %d", len(articles[1].ArticleNames))
	}
	if articles[1].ArticleNames[0].Name != "Shirt" {
		t.Errorf("Expected Shirt, get %s", articles[1].ArticleNames[0].Name)
	}
	if articles[1].ArticleNames[1].Name != "Maglietta" {
		t.Errorf("Expected Maglietta, get %s", articles[1].ArticleNames[1].Name)
	}

	// Do it again

	articles = []*Article{
		{
			Code: "123",
			ArticleNames: []*ArticleName{
				{Name: "Jacket", Language: "en"},
				{Name: "Giacca", Language: "it"},
			},
		},
		{
			Code: "456",
			ArticleNames: []*ArticleName{
				{Name: "Shirt", Language: "en"},
				{Name: "Maglietta", Language: "it"},
			},
		},
	}

	err = DB.Clauses(clause.OnConflict{UpdateAll: true}).
		Session(&gorm.Session{CreateBatchSize: 1000, FullSaveAssociations: true}).
		Create(&articles).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	err = DB.Preload("ArticleNames").Find(&articles).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if len(articles) != 2 {
		t.Errorf("Expected 2 articles, found %d", len(articles))
	}
	if articles[0].Code != "123" {
		t.Errorf("Expected article 1 with code 123, found %s", articles[0].Code)
	}
	if len(articles[0].ArticleNames) != 2 {
		t.Errorf("Expected article 1 with 2 names, found %d", len(articles[0].ArticleNames))
	}
	if articles[0].ArticleNames[0].Name != "Jacket" {
		t.Errorf("Expected Jacket, get %s", articles[0].ArticleNames[0].Name)
	}
	if articles[0].ArticleNames[1].Name != "Giacca" {
		t.Errorf("Expected Giacca, get %s", articles[0].ArticleNames[1].Name)
	}
	if articles[1].Code != "456" {
		t.Errorf("Expected article 2 with code 456, found %s", articles[1].Code)
	}
	if len(articles[1].ArticleNames) != 2 {
		t.Errorf("Expected article 2 with 2 names, found %d", len(articles[1].ArticleNames))
	}
	if articles[1].ArticleNames[0].Name != "Shirt" {
		t.Errorf("Expected Shirt, get %s", articles[1].ArticleNames[0].Name)
	}
	if articles[1].ArticleNames[1].Name != "Maglietta" {
		t.Errorf("Expected Maglietta, get %s", articles[1].ArticleNames[1].Name)
	}

	// Do it again with updated values

	articles = []*Article{
		{
			Code: "123",
			ArticleNames: []*ArticleName{
				{Name: "Jacket", Language: "en"},
				{Name: "Giacca", Language: "it"},
			},
		},
		{
			Code: "456bis",
			ArticleNames: []*ArticleName{
				{Name: "T-Shirt", Language: "en"},
				{Name: "Maglietta", Language: "it"},
			},
		},
	}

	err = DB.Clauses(clause.OnConflict{UpdateAll: true}).
		Session(&gorm.Session{CreateBatchSize: 1000, FullSaveAssociations: true}).
		Create(&articles).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	err = DB.Preload("ArticleNames").Find(&articles).Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if len(articles) != 2 {
		t.Errorf("Expected 2 articles, found %d", len(articles))
	}
	if articles[0].Code != "123" {
		t.Errorf("Expected article 1 with code 123, found %s", articles[0].Code)
	}
	if len(articles[0].ArticleNames) != 2 {
		t.Errorf("Expected article 1 with 2 names, found %d", len(articles[0].ArticleNames))
	}
	if articles[0].ArticleNames[0].Name != "Jacket" {
		t.Errorf("Expected Jacket, get %s", articles[0].ArticleNames[0].Name)
	}
	if articles[0].ArticleNames[1].Name != "Giacca" {
		t.Errorf("Expected Giacca, get %s", articles[0].ArticleNames[1].Name)
	}
	if articles[1].Code != "456bis" {
		t.Errorf("Expected article 2 with code 456bis, found %s", articles[1].Code)
	}
	if len(articles[1].ArticleNames) != 3 {
		t.Errorf("Expected article 2 with 3 names, found %d", len(articles[1].ArticleNames))
	} else {
		if articles[1].ArticleNames[0].Name != "Shirt" {
			t.Errorf("Expected Shirt, get %s", articles[1].ArticleNames[0].Name)
		}
		if articles[1].ArticleNames[1].Name != "Maglietta" {
			t.Errorf("Expected Maglietta, get %s", articles[1].ArticleNames[1].Name)
		}
		if articles[1].ArticleNames[2].Name != "T-Shirt" {
			t.Errorf("Expected T-Shirt, get %s", articles[1].ArticleNames[2].Name)
		}
	}
}
