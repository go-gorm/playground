package main

import (
	"errors"
	"gorm.io/gorm"
	"sync"
)

var form = &formService{
	mutex: &sync.Mutex{},
}

type formService struct {
	mutex *sync.Mutex
}

type totalData struct {
	Channel   `json:"channel"`
	Worker    uint `json:"worker"`
	Mechanism uint `json:"mechanism"`
	Sign      uint `json:"sign"`
}

func (srv *formService) DataWithCount(fields interface{}, maps *Form, limit, offset int) (info []*totalData, count int64, err error) {
	query := DB.Model(&Form{}).Debug().
		Select(fields).
		Joins("left join user on user.id = form.user_id").
		Joins("left join channel on channel.id = user.channel_id").
		Where(maps).
		Group("channel.id").
		Count(&count).
		Offset(offset).
		Limit(limit).
		Find(&info)
	if !errors.Is(query.Error, gorm.ErrRecordNotFound) && query.Error != nil {
		panic(query.Error)
	}
	return
}

func (srv *formService) DataWithoutCount(maps *Form, limit, offset int) (info []*totalData, count int64, err error) {
	query := DB.Model(&Form{}).Debug().
		Select(
			//"channel.*",
			"sum(CASE WHEN form.user_type=1 THEN 1 ELSE 0 END) as worker",
			"sum(CASE WHEN form.user_type=2 THEN 1 ELSE 0 END) as mechanism",
			"sum(CASE WHEN user.sign_time is not null THEN 1 ELSE 0 END) as sign").
		Joins("left join user on user.id = form.user_id").
		Joins("left join channel on channel.id = user.channel_id").
		Where(maps).
		Group("channel.id").
		//Count(&count).
		Offset(offset).
		Limit(limit).
		Find(&info)
	if !errors.Is(query.Error, gorm.ErrRecordNotFound) && query.Error != nil {
		panic(query.Error)
	}
	return
}
