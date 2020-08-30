package v4

import (
	"gorm.io/gorm"
)

func Migrate_17_tipo_decimal_exoneraciones_otros_cargos(tx *gorm.DB, m gorm.Migrator) error {
	if err := m.AlterColumn(&Exoneracion{}, "porcentaje_exoneracion"); err != nil {
		return err
	}

	if err := m.AlterColumn(&OtrosCargos{}, "porcentaje"); err != nil {
		return err
	}
	return nil
}
