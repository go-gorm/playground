package v8migration

import "gorm.io/gorm"

var (
	Migrate_8_codigos = func(tx *gorm.DB, m gorm.Migrator) error {
		err := m.AutoMigrate(&Codigo{}, &Grupo{})
		if err != nil {
			return err
		}
		return nil
	}
)
