package model42

import (
	"gorm.io/gorm"
	"gorm.io/playground/pkg/model/migrator"
	v1migration "gorm.io/playground/pkg/model/v42/migration/v1"
	v10migration "gorm.io/playground/pkg/model/v42/migration/v10"
	v11migration "gorm.io/playground/pkg/model/v42/migration/v11"
	v12migration "gorm.io/playground/pkg/model/v42/migration/v12"
	v13migration "gorm.io/playground/pkg/model/v42/migration/v13"
	v3migration "gorm.io/playground/pkg/model/v42/migration/v3"
	v4migration "gorm.io/playground/pkg/model/v42/migration/v4"
	v5migration "gorm.io/playground/pkg/model/v42/migration/v5"
	v6migration "gorm.io/playground/pkg/model/v42/migration/v6"
	v7migration "gorm.io/playground/pkg/model/v42/migration/v7"
	v8migration "gorm.io/playground/pkg/model/v42/migration/v8"
	v9migration "gorm.io/playground/pkg/model/v42/migration/v9"
)

var (
	return_nil = func(tx *gorm.DB, m gorm.Migrator) error {
		return nil
	}
)

var Migrations = []*migrator.SingleMigration{
	{
		ID:       "1-inicial",
		Migrate:  v1migration.Migrate_1_inicial,
		Rollback: v1migration.Rollback_1_inicial,
	},
	{
		ID:       "3-campo-tipo",
		Migrate:  v3migration.Migrate_3_campo_tipo,
		Rollback: v3migration.Rollback_migrate_3_campo_tipo,
	},
	{
		ID:       "4-eliminar-emisor",
		Migrate:  v4migration.Migrate_4_eliminar_emisor,
		Rollback: v4migration.Rollback_4_eliminar_emisor,
	},
	{
		ID:       "5-agregar-estado-documento",
		Migrate:  v5migration.Migrate_5_agregar_estado_documento,
		Rollback: v5migration.Rollback_5_agregar_estado_documento,
	},
	{
		ID:       "6-agregar-locacion-documento",
		Migrate:  v6migration.Migrate_6_agregar_locacion_documento,
		Rollback: return_nil,
	},
	{
		ID:       "7-agregar-secuenciales-grupo",
		Migrate:  v7migration.Migrate_7_agregar_secuenciales_grupo,
		Rollback: return_nil,
	},
	{
		ID:       "8-codigos",
		Migrate:  v8migration.Migrate_8_codigos,
		Rollback: return_nil,
	},
	{
		ID:       "9-cambia nombre de campo en tabla",
		Migrate:  v9migration.Migrate_9_cambia_nombre_de_campo_en_tabla,
		Rollback: v9migration.Rollback_9_cambia_nombre_de_campo_en_tabla,
	},
	{
		ID:       "10-agrega intentos de envio de comprobante",
		Migrate:  v10migration.Migrate_10_agrega_intentos_de_envio_de_comprobante,
		Rollback: return_nil,
	},
	{
		ID:       "11-tabla recepcion",
		Migrate:  v11migration.Migrate_11_tabla_recepcion,
		Rollback: return_nil,
	},
	{
		ID:       "12-Tabla de consecutivos",
		Migrate:  v12migration.Migrate_12_Tabla_de_consecutivos,
		Rollback: return_nil,
	},
	{
		ID:       "13-Agrega campo NumeroConsecutivoReceptor a MensajeHacienda",
		Migrate:  v13migration.Migrate_13_Agrega_campo_NumeroConsecutivoReceptor_a_MensajeHacienda,
		Rollback: return_nil,
	},
}
