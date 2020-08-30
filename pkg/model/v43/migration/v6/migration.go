package v6

import "gorm.io/gorm"

func Migrate_19_agregar_pagado_el(tx *gorm.DB, m gorm.Migrator) error {
	return tx.AutoMigrate(&Suscripcion{})
}
