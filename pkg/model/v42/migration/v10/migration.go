package v10migration

import "gorm.io/gorm"

var (
	Migrate_10_agrega_intentos_de_envio_de_comprobante = func(tx *gorm.DB, m gorm.Migrator) error {
		return tx.AutoMigrate(Documento{})
	}
)
