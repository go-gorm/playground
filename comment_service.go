package main

import (
	"errors"
	"gorm.io/gorm"
	"sync"
)

type commentService struct {
	mutex *sync.Mutex
	table string
}

func (srv *commentService) HaveCountList(maps *Comment, where interface{}, limit, offset int) (info []*Comment, count int64, err error) {
	query := DB.Model(&Comment{}).
		Unscoped().
		Joins("User").
		Where(where).Where(maps).
		Group(srv.table + ".id").
		Count(&count).Offset(offset).Limit(limit).
		Order(srv.table + ".id desc").
		Find(&info)
	if err = query.Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	}
	return
}

func (srv *commentService) NormalList(maps *Comment, where interface{}, limit, offset int) (info []*Comment, count int64, err error) {
	query := DB.Model(&Comment{}).
		Unscoped().
		Joins("User").
		Where(where).Where(maps).
		Group(srv.table + ".id").
		Offset(offset).Limit(limit).
		Order(srv.table + ".id desc").
		Find(&info)
	if err = query.Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	}
	return
}
