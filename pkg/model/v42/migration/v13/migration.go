package v13migration

import "gorm.io/gorm"

var (
	Migrate_13_Agrega_campo_NumeroConsecutivoReceptor_a_MensajeHacienda = func(tx *gorm.DB, m gorm.Migrator) error {
		return tx.AutoMigrate(MensajeHacienda{})
	}
)
