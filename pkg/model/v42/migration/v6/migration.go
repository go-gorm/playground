package v6migration

import "gorm.io/gorm"

var (
	Migrate_6_agregar_locacion_documento = func(tx *gorm.DB, m gorm.Migrator) error {
		type Documento struct {
			Locacion string
		}
		return m.AutoMigrate(&Documento{})
	}
)
