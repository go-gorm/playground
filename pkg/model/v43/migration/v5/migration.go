package v5

import "gorm.io/gorm"

func Migrate_18_invalido_desde(tx *gorm.DB, m gorm.Migrator) error {
	return tx.AutoMigrate(&Credencial{})
}
