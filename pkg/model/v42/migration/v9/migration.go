package v9migration

import "gorm.io/gorm"

var (
	Migrate_9_cambia_nombre_de_campo_en_tabla = func(tx *gorm.DB, m gorm.Migrator) error {
		type LineaDetalle struct {
			gorm.Model
			CodigoCodigo string
		}
		return m.RenameColumn(&LineaDetalle{}, "codigo_codigo", "codigo")
	}
	Rollback_9_cambia_nombre_de_campo_en_tabla = func(tx *gorm.DB, m gorm.Migrator) error {
		type LineaDetalle struct {
			gorm.Model
			Codigo string
		}
		return m.RenameColumn(&LineaDetalle{}, "codigo", "codigo_codigo")
	}
)
