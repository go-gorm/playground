package main

import "gorm.io/gorm"

type Distribution struct {
	ID          []byte           `gorm:"type:bytea;primaryKey;uniqueIndex;not null"`
	IdentityKey string           `gorm:"type:text;default:'';not null"`
	Description string           `gorm:"type:text"`
	Places      []Place          `gorm:"many2many:distribution_places;"`
}

type Place struct {
	gorm.Model
	Distributions    []Distribution `gorm:"many2many:distribution_places;"`
	IdentityKey string `gorm:"uniqueIndex:places_merchant_identity"`
}
