package main

import (
	"encoding/json"
	"gorm.io/gorm/clause"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	tag1 := Tag{
		Name: "tag1",
	}
	tag2 := Tag{
		Name: "tag2",
	}
	tag3 := Tag{
		Name: "tag3",
	}

	category := Category{
		Name: "cate1",
	}

	DB.Create(&tag1)
	DB.Create(&tag2)
	DB.Create(&tag3)
	DB.Create(&category)

	post1 := Post{
		Title:    "post1",
		Tags:     []Tag{tag1, tag2},
		Category: category,
	}
	DB.Create(&post1)

	var post Post
	DB.Where("id = ?", 1).Find(&post)
	DB.Model(&post).Association("Tags").Append(&tag3)
	DB.Preload(clause.Associations).Find(&post)

	res1, _ := json.Marshal(post.Tags)
	tags := make([]Tag, 0)
	DB.Model(&post).Association("Tags").Find(&tags)
	res2, _ := json.Marshal(tags)
	if string(res1) != string(res2) {
		t.Errorf("失败, res1: %s ,res2: %s", res1, res2)
	}
}

type Post struct {
	ID         uint64
	Title      string
	Tags       []Tag `gorm:"many2many:post_tags"`
	CategoryID uint64
	Category   Category
}

type Category struct {
	ID   uint64
	Name string
}

type Tag struct {
	ID   uint64
	Name string
}
