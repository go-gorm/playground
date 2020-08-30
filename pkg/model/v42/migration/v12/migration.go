package v12migration

import (
	"gorm.io/gorm"
	"gorm.io/playground/pkg/model"
)

var (
	Migrate_12_Tabla_de_consecutivos = func(tx *gorm.DB, m gorm.Migrator) error {
		err := tx.AutoMigrate(&Consecutivo{})
		if err != nil {
			return err
		}

		// migra los contadores de los grupos a la
		// nueva tabla
		rows, err := tx.Model(&Grupo{}).Order("created_at").Rows()
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			var grupo Grupo
			errdb := tx.ScanRows(rows, &grupo)
			if errdb != nil {
				return errdb
			}
			if grupo.NumeroDocumento > 0 {
				fe := Consecutivo{
					GrupoID:  grupo.ID,
					Tipo:     model.CFactura,
					Sucursal: 1,
					Caja:     1,
					Contador: grupo.NumeroDocumento,
				}
				tx.Create(&fe)
			}
			if grupo.NumeroDocumentoNC > 0 {
				nc := Consecutivo{
					GrupoID:  grupo.ID,
					Tipo:     model.CCredito,
					Sucursal: 1,
					Caja:     1,
					Contador: grupo.NumeroDocumentoNC,
				}
				tx.Create(&nc)
			}
			if grupo.NumeroDocumentoND > 0 {
				nd := Consecutivo{
					GrupoID:  grupo.ID,
					Tipo:     model.CDebito,
					Sucursal: 1,
					Caja:     1,
					Contador: grupo.NumeroDocumentoND,
				}
				tx.Create(&nd)
			}
			if grupo.NumeroDocumentoTC > 0 {
				te := Consecutivo{
					GrupoID:  grupo.ID,
					Tipo:     model.CTiquete,
					Sucursal: 1,
					Caja:     1,
					Contador: grupo.NumeroDocumentoTC,
				}
				tx.Create(&te)
			}
		}
		// Borra las columnas
		m.DropColumn(&Grupo{}, "numero_documento")
		m.DropColumn(&Grupo{}, "numero_documento_nc")
		m.DropColumn(&Grupo{}, "numero_documento_nd")
		m.DropColumn(&Grupo{}, "numero_documento_tc")
		return nil
	}
)
