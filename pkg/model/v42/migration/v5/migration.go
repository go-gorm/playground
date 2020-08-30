package v5migration

import "gorm.io/gorm"

var (
	Migrate_5_agregar_estado_documento = func(tx *gorm.DB, m gorm.Migrator) error {
		// Crea las nuevas columnas
		return m.AddColumn(&Documento{}, "estado")
	}
	Rollback_5_agregar_estado_documento = func(tx *gorm.DB, m gorm.Migrator) error {
		// Borra las viejas columnas
		return m.DropColumn(&Documento{}, "estado")
	}
)
