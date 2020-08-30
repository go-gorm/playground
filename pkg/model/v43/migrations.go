package models43

import (
	"gorm.io/gorm"
	"gorm.io/playground/pkg/model/migrator"
	v1 "gorm.io/playground/pkg/model/v43/migration/v1"
	v2 "gorm.io/playground/pkg/model/v43/migration/v2"
	v3 "gorm.io/playground/pkg/model/v43/migration/v3"
	v4 "gorm.io/playground/pkg/model/v43/migration/v4"
	v5 "gorm.io/playground/pkg/model/v43/migration/v5"
	v6 "gorm.io/playground/pkg/model/v43/migration/v6"
	v7 "gorm.io/playground/pkg/model/v43/migration/v7"
	v8 "gorm.io/playground/pkg/model/v43/migration/v8"
	v9 "gorm.io/playground/pkg/model/v43/migration/v9"
)

var (
	return_nil = func(tx *gorm.DB, m gorm.Migrator) error {
		return nil
	}
)

var Migrations = []*migrator.SingleMigration{
	{
		ID:       "14-prepara la base de datos para comprobantes electrónicos 4.3",
		Migrate:  v1.Migrate_14_Actualiza_tablas_para_version_4_3,
		Rollback: return_nil,
	},
	{
		ID:       "15-Crea nueva tabla de emisor y receptor y copia datos del grupo en cada comprobante a tabla de emisores",
		Migrate:  v2.Migrate_15_Crea_nueva_tabla_de_Emisor_y_Receptor,
		Rollback: return_nil,
	},
	{
		ID:       "15-Crea todas las actividades economicas dentro de la tabla",
		Migrate:  v2.Migrar_15_Crea_Actividades,
		Rollback: return_nil,
	},
	{
		ID:       "16-ubicaciones codificadas",
		Migrate:  v3.Migrate_16_ubicaciones,
		Rollback: v3.Rollback_16_ubicaciones,
	},
	{
		ID:       "17-cambia precision decimal en Exoneraciones y Otros Cargos",
		Migrate:  v4.Migrate_17_tipo_decimal_exoneraciones_otros_cargos,
		Rollback: return_nil,
	},
	{
		ID:       "18-indica desde cuando los credenciales son inválidos",
		Migrate:  v5.Migrate_18_invalido_desde,
		Rollback: return_nil,
	},
	{
		ID:       "19-crea campo de ultimo pago de suscripción hecho",
		Migrate:  v6.Migrate_19_agregar_pagado_el,
		Rollback: return_nil,
	},
	{
		ID:       "20-mueve permisos de suscripción a usuarios",
		Migrate:  v7.Migrate_20_crear_permisos_y_asocacion_con_usuarios,
		Rollback: return_nil,
	},
	{
		ID:       "21-Elimina columnas de suscripcion",
		Migrate:  v8.Migrate,
		Rollback: return_nil,
	},
	{
		ID:       "22-Renombra las tablas al nuevo esquema de gorm v2",
		Migrate:  v9.Migrate,
		Rollback: return_nil,
	},
}
