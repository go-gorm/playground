package main

import (
	"context"
	"gorm.io/sharding"
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type ShardingModel struct {
	ID  uint `gorm:"column:id;type:int unsigned;not null;auto_increment;primary key;'" json:"id"`
	UID uint `gorm:"column:uid;type:int unsigned;not null;" json:"uid"`
}

func (*ShardingModel) TableName() string {
	return "sharding_model"
}

func TestGORM(t *testing.T) {
	err := DB.Use(sharding.Register(sharding.Config{
		ShardingKey:         "uid",
		NumberOfShards:      2,
		PrimaryKeyGenerator: sharding.PKCustom,
		PrimaryKeyGeneratorFn: func(tableIdx int64) int64 {
			return 0
		},
	}, &ShardingModel{}))
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
		return
	}

	err = DB.AutoMigrate(&ShardingModel{})
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan error)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := DB.Model(&ShardingModel{}).Where("uid = ?", 100).Find(&[]ShardingModel{}).Error
				if err != nil {
					ch<-err
					return
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := DB.Model(&ShardingModel{}).Where("uid = ?", 100).Find(&[]ShardingModel{}).Error
				if err != nil {
					ch<-err
					return
				}
			}
		}
	}()

	select {
	case <-time.After(time.Millisecond * 100):
		cancel()
	case err = <-ch:
		cancel()
		if err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
	}
}
