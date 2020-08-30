package v8

import "gorm.io/gorm"

func Migrate(tx *gorm.DB, m gorm.Migrator) error {
	m.DropColumn(&Suscripcion{}, "identificacion_tipo")
	m.DropColumn(&Suscripcion{}, "identificacion_numero")
	m.DropColumn(&Suscripcion{}, "nombre")

	err := m.RenameTable("suscripcions", "suscripciones")

	return err
}
