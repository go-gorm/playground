package v9

import "gorm.io/gorm"

func Migrate(tx *gorm.DB, m gorm.Migrator) error {
	if m.HasTable(actividad) {
		err := m.RenameTable(actividad, &Actividad{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(barrio) {
		err := m.RenameTable(barrio, &Barrio{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(canton) {
		err := m.RenameTable(canton, &Canton{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(codigo) {
		err := m.RenameTable(codigo, &Codigo{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(codigo_comercial) {
		err := m.RenameTable(codigo_comercial, &CodigoComercial{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(consecutivo) {
		err := m.RenameTable(consecutivo, &Consecutivo{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(credencial) {
		err := m.RenameTable(credencial, &Credencial{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(distrito) {
		err := m.RenameTable(distrito, &Distrito{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(documento) {
		err := m.RenameTable(documento, &Documento{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(emisor) {
		err := m.RenameTable(emisor, &Emisor{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(exoneracion) {
		err := m.RenameTable(exoneracion, &Exoneracion{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(grupo) {
		err := m.RenameTable(grupo, &Grupo{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(ga) {
		err := m.RenameTable(ga, &Ga{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(usuario_permisos) {
		err := m.RenameTable(usuario_permisos, &UsuarioPermisos{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(impuesto) {
		err := m.RenameTable(impuesto, &Impuesto{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(informacion_referencia) {
		err := m.RenameTable(informacion_referencia, &InformacionReferencia{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(linea_detalle) {
		err := m.RenameTable(linea_detalle, &LineaDetalle{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(mensaje_hacienda) {
		err := m.RenameTable(mensaje_hacienda, &MensajeHacienda{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(o_auth_response) {
		err := m.RenameTable(o_auth_response, &OAuthResponse{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(otros_cargos) {
		err := m.RenameTable(otros_cargos, &OtrosCargos{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(permiso) {
		err := m.RenameTable(permiso, &Permiso{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(recepcion) {
		err := m.RenameTable(recepcion, &Recepcion{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(receptor) {
		err := m.RenameTable(receptor, &Receptor{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(suscripcion) {
		err := m.RenameTable(suscripcion, &Suscripcion{})
		if err != nil {
			return err
		}
	}

	if m.HasTable(usuario) {
		err := m.RenameTable(usuario, &Usuario{})
		if err != nil {
			return err
		}
	}

	return nil
}
