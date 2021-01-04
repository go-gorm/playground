package main

import (
	"errors"
	"gorm.io/gorm"
	"sync"
)

var article = &articleService{
	mutex: &sync.Mutex{},
}

type articleService struct {
	mutex *sync.Mutex
}

func (srv *articleService) List(maps *Article, where interface{}, limit, offset int) (info []Article, count int64, err error) {
	query := DB.Model(&Article{}).
		Where(where).
		Where(maps).
		Count(&count).
		Offset(offset).
		Limit(limit).
		Order("id desc").Find(&info)
	if err = query.Error; !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		panic(query.Error)
	}
	return
}
