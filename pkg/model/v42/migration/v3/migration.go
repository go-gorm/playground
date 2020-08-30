package v3migration

import "gorm.io/gorm"

var (
	Migrate_3_campo_tipo = func(tx *gorm.DB, m gorm.Migrator) error {
		return tx.Exec("ALTER TABLE t_impuesto ALTER COLUMN codigo TYPE integer USING (codigo::integer)").Error
	}
	Rollback_migrate_3_campo_tipo = func(tx *gorm.DB, m gorm.Migrator) error {
		return tx.Exec("ALTER TABLE t_impuesto ALTER COLUMN codigo TYPE text USING (codigo::text)").Error
	}
)
