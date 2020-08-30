package v7migration

import "gorm.io/gorm"

var (
	Migrate_7_agregar_secuenciales_grupo = func(tx *gorm.DB, m gorm.Migrator) error {
		return tx.AutoMigrate(&Grupo{})
	}
)
