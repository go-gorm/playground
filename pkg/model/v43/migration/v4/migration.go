package v4

import (
	"gorm.io/gorm"
)

func Migrate_17_tipo_decimal_exoneraciones_otros_cargos(tx *gorm.DB, m gorm.Migrator) error {
	err := tx.Exec("ALTER TABLE t_exoneracion ALTER COLUMN porcentaje_exoneracion TYPE DECIMAL(4,2);").Error
	if err != nil {
		return err
	}
	err = tx.Exec("ALTER TABLE t_otros_cargos ALTER COLUMN porcentaje TYPE DECIMAL(4,2);").Error
	return err
}
