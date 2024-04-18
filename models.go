package main

import (
	"time"

	"gorm.io/gorm"
)

type (
	Area struct {
		ID        uint      `gorm:"primaryKey"`
		CreatedAt time.Time `gorm:"autoCreateTime:nano"`
		UpdatedAt time.Time `gorm:"autoUpdateTime:nano"`
		DeletedAt gorm.DeletedAt
	}
	City struct {
		ID       uint `gorm:"primaryKey"`
		RegionID uint
		Region   Region `gorm:"foreignKey:RegionID;references:ID"`
		AreaID   uint
		Area     Area `gorm:"foreignKey:AreaID;references:ID"`
	}
	Region struct {
		ID             uint `gorm:"primaryKey"`
		CountryID      uint
		Country        Country `gorm:"foreignKey:CountryID;references:ID"`
		HappiestCityID uint
		// Commenting out the following line will not cause error, but foreign key is not created
		City *City `gorm:"foreignKey:HappiestCityID;references:ID"`
	}
	Country struct {
		ID uint `gorm:"primaryKey"`
	}
)
