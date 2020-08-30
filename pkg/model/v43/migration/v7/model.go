package v7

import (
	"gorm.io/gorm"
	"github.com/lib/pq"
)

type Permiso struct {
	gorm.Model
	Sujeto string `gorm:"unique_index:idx_sujeto_accion"`
	Accion string `gorm:"unique_index:idx_sujeto_accion"`
}

type Usuario struct {
	gorm.Model
	SuscripcionID int       `json:"suscripcion_id"`
	Permisos      []Permiso `gorm:"many2many:usuario_permisos;"`
}

type Suscripcion struct {
	gorm.Model
	Usuarios []Usuario      `json:"usuarios"`
	Permisos pq.StringArray `gorm:"type:varchar(64)[]" json:"permisos"`
}
