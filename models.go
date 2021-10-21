package main

import "gorm.io/gorm"

type Layer struct {
	gorm.Model
	Labels []*Label
	Points []*Point
}

type Label struct {
	gorm.Model
	Name string

	LayerID int
	Layer   *Layer

	ParentID *int
	Children []*Label `gorm:"foreignkey:ParentID"`
}

type Point struct {
	gorm.Model

	LayerID int
	Layer   *Layer

	Labels []*Label `gorm:"many2many:points_labels;"`
}
