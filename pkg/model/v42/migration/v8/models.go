package v8migration

import "gorm.io/gorm"

type Codigo struct {
	gorm.Model
	Codigo  string `gorm:"not null;size:20;uniqueIndex:idx_codigo_grupoid"`
	Detalle string `gorm:"not null;size:160"`
	GrupoID uint   `gorm:"uniqueIndex:idx_codigo_grupoid"`
}

type Grupo struct {
	gorm.Model
	Codigos []Codigo
}
