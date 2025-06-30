package main

import (
	"gorm.io/plugin/soft_delete"
)

// --- Models ---
type Dataset struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"comment:Name;size:30;index;not null;"`
	Columns []*DatasetColumn
}

type DatasetColumn struct {
	ID        uint                  `gorm:"primaryKey"`
	DatasetID uint
	Dataset   Dataset
	DeletedAt soft_delete.DeletedAt `gorm:"comment:删除时间;uniqueIndex:idx_dc_ddn;not null;"`
	Name      string                `gorm:"comment:名称;size:30;uniqueIndex:idx_dc_ddn;not null;"`
	IsSystem  bool                  `gorm:"comment:是否系统列;not null;default:false;"`
	IsHidden  bool                  `gorm:"comment:是否隐藏;not null;"`
}