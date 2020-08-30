package v5migration

import "gorm.io/gorm"

var (
	Migrate_5_agregar_estado_documento = func(tx *gorm.DB, m gorm.Migrator) error {
		// Crea las nuevas columnas
		if err := m.AddColumn(&Documento{}, "estado"); err != nil {
			return err
		}

		if err := m.AddColumn(&Documento{}, "locacion"); err != nil {
			return err
		}
	}
	Rollback_5_agregar_estado_documento = func(tx *gorm.DB, m gorm.Migrator) error {
		// Borra las viejas columnas
		m.DropColumn(&Documento{}, "estado")
		m.DropColumn(&Documento{}, "locacion")
		return nil
	}
)
