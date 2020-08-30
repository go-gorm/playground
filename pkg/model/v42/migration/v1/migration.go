package v1migration

import "gorm.io/gorm"

var (
	Migrate_1_inicial = func(tx *gorm.DB, m gorm.Migrator) error {
		return tx.AutoMigrate(
			&Suscripcion{},
			&Usuario{},
			&Grupo{},
			&Documento{},
			&MensajeHacienda{},
			&Emisor{},
			&Impuesto{},
			&InformacionReferencia{},
			&LineaDetalle{},
			&Receptor{},
			&OAuthResponse{},
			&Credencial{},
		)
	}
	Rollback_1_inicial = func(tx *gorm.DB, m gorm.Migrator) error {
		return m.DropTable(
			&Emisor{},
			&Documento{},
			&Impuesto{},
			&InformacionReferencia{},
			&LineaDetalle{},
			&MensajeHacienda{},
			&Receptor{},
			&Usuario{},
			&OAuthResponse{},
			&Credencial{},
			&Grupo{},
			&Suscripcion{})
	}
)
