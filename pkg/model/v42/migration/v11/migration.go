package v11migration

import "gorm.io/gorm"

var (
	Migrate_11_tabla_recepcion = func(tx *gorm.DB, m gorm.Migrator) error {
		return m.AutoMigrate(&Recepcion{})
	}
)
