package v8

import "gorm.io/gorm"

func Migrate(tx *gorm.DB, m gorm.Migrator) error {
	m.DropColumn(&Suscripcion{}, "identificacion_tipo")
	m.DropColumn(&Suscripcion{}, "identificacion_numero")
	m.DropColumn(&Suscripcion{}, "nombre")

	if m.HasTable("suscripcions") {
		err := m.RenameTable("suscripcions", "suscripciones")
		if err != nil {
			return err
		}
	}

	return nil
}
