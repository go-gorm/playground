package v1

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	model42 "gorm.io/playground/pkg/model/v42"
)

var (
	Migrate_14_Actualiza_tablas_para_version_4_3 = func(tx *gorm.DB, m gorm.Migrator) error {
		// recrea las tablas o añade campos faltantes
		if err := m.AutoMigrate(&OtrosCargos{}); err != nil {
			return err
		}
		if err := m.AutoMigrate(&Documento{}); err != nil {
			return err
		}
		if err := m.AutoMigrate(&LineaDetalle{}); err != nil {
			return err
		}
		if err := m.AutoMigrate(&Exoneracion{}); err != nil {
			return err
		}
		if err := m.AutoMigrate(&Impuesto{}); err != nil {
			return err
		}
		if err := m.AutoMigrate(&Receptor{}); err != nil {
			return err
		}
		if err := m.AutoMigrate(&Grupo{}); err != nil {
			return err
		}
		if err := m.AutoMigrate(&CodigoComercial{}); err != nil {
			return err
		}
		// ajusta el tamaño de algunos campos en algunas tablas
		err := m.AlterColumn(&Grupo{}, "nombre")
		if err != nil {
			return err
		}

		err = m.AlterColumn(&Grupo{}, "otras_senas")
		if err != nil {
			return err
		}

		// Migra todos los códigos de productos en cada
		// linea de detalle a su nueva tabla
		rows, err := tx.Model(&model42.LineaDetalle{}).Order("created_at").Rows()
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			var linea model42.LineaDetalle
			bErr := tx.ScanRows(rows, &linea)
			if bErr != nil {
				return bErr
			}
			bErr = tx.Where("id = ?", linea.ID).First(&linea).Error
			if bErr != nil {
				return bErr
			}

			nuevalinea := LineaDetalle{
				Model: gorm.Model{
					ID:        linea.ID,
					CreatedAt: linea.CreatedAt,
				},
				NumeroLinea:           linea.NumeroLinea,
				Cantidad:              linea.Cantidad,
				UnidadMedida:          linea.UnidadMedida,
				UnidadMedidaComercial: linea.UnidadMedidaComercial,
				Detalle:               linea.Detalle,
				PrecioUnitario:        linea.PrecioUnitario,
				MontoTotal:            linea.MontoTotal,
				MontoDescuento:        linea.MontoDescuento,
				NaturalezaDescuento:   linea.NaturalezaDescuento,
				SubTotal:              linea.SubTotal,
				MontoTotalLinea:       linea.MontoTotalLinea,
				DocumentoID:           linea.DocumentoID,
			}

			// borra la linea de la tabla antes de insertar
			tx.Unscoped().Delete(&linea)

			bErr = tx.Create(&nuevalinea).Error
			if bErr != nil {
				return bErr
			}

			codigocomercial := CodigoComercial{
				Tipo:           linea.CodigoTipo,
				Codigo:         linea.Codigo,
				LineaDetalleID: linea.ID,
			}
			bErr = tx.Create(&codigocomercial).Error
			if bErr != nil {
				return bErr
			}
			// mueve los impuestos
			for _, impuesto := range linea.Impuesto {
				if impuesto.Tarifa.Equal(decimal.NewFromFloat(13)) {
					nuevo_impuesto := Impuesto{
						Model: gorm.Model{
							CreatedAt: impuesto.CreatedAt,
							ID:        impuesto.ID,
						},
						Codigo:         1,
						CodigoTarifa:   8,
						Tarifa:         decimal.NewFromFloat(13),
						Monto:          impuesto.Monto,
						LineaDetalleID: impuesto.LineaDetalleID,
					}
					iErr := tx.Create(&nuevo_impuesto).Error
					if iErr != nil {
						return iErr
					}
				} else if impuesto.Tarifa.Equal(decimal.NewFromFloat(10)) {
					// los cargos del
					// impuesto de servicio
					// del 10% van a nivel
					// del comprobante
					otros_cagos := OtrosCargos{
						TipoDocumento: 6,
						Detalle:       "Impuesto de servicio del 10%",
						MontoCargo:    impuesto.Monto,
						Porcentaje:    decimal.NewFromFloat(10),
						DocumentoID:   linea.DocumentoID,
					}
					oErr := tx.Create(&otros_cagos).Error
					if oErr != nil {
						return oErr
					}
				}
			}
		}
		return nil
	}
)
