package v4migration

import "gorm.io/gorm"

var (
	Migrate_4_eliminar_emisor = func(tx *gorm.DB, m gorm.Migrator) error {
		// Borra la tabla Emisor
		_ = m.DropTable(&Emisor{})

		// Crea las nuevas columnas
		if err := tx.AutoMigrate(&MensajeHacienda{}, &Documento{}); err != nil {
			return err
		}
		// Borra las viejas columnas
		_ = m.DropColumn(&Documento{}, "emisor_id")
		return nil
	}
	Rollback_4_eliminar_emisor = func(tx *gorm.DB, m gorm.Migrator) error {
		// Crea la tabla y las columnas
		if err := tx.AutoMigrate(&Emisor{}, &Documento{}); err != nil {
			return err
		}
		// Borra las viejas columnas
		if err := m.DropColumn(&MensajeHacienda{}, "grupo_id"); err != nil {
			return err
		}
		if err := m.DropColumn(&Documento{}, "grupo_id"); err != nil {
			return err
		}
		return nil
	}
)
