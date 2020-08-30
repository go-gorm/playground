package v2

import (
	"encoding/json"
	"os"

	"gorm.io/gorm"
	model42 "gorm.io/playground/pkg/model/v42"
)

type ActividadesEconomicas struct {
	Actividades []Actividad `json:"actividades"`
}

var (
	Migrate_15_Crea_nueva_tabla_de_Emisor_y_Receptor = func(tx *gorm.DB, m gorm.Migrator) error {
		err := tx.AutoMigrate(&Grupo{}, &Emisor{}, &Documento{}, &Actividad{})
		if err != nil {
			return err
		}

		// copia la información del grupo que emitió el documento a la
		// nueva tabla Emisor conservando la asociación del comprobante
		rows, err := tx.Model(&model42.Documento{}).Order("created_at").Rows()
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			var doc model42.Documento
			errdb := tx.ScanRows(rows, &doc)
			if errdb != nil {
				return errdb
			}
			// carga los datos para el campo de Grupo
			gErr := tx.Where("id = ?", doc.GrupoID).First(&doc.Grupo).Error
			if gErr != nil {
				return gErr
			}
			emisor := Emisor{
				Nombre:               doc.Grupo.Nombre,
				IdentificacionNumero: doc.Grupo.IdentificacionNumero,
				GrupoBase: GrupoBase{
					IdentificacionTipo: doc.Grupo.IdentificacionTipo,
					Provincia:          doc.Grupo.Provincia,
					Canton:             doc.Grupo.Canton,
					Distrito:           doc.Grupo.Distrito,
					Barrio:             doc.Grupo.Barrio,
					OtrasSenas:         doc.Grupo.OtrasSenas,
					TelCodigoPais:      doc.Grupo.TelCodigoPais,
					TelNumTelefono:     doc.Grupo.TelNumTelefono,
					FaxCodigoPais:      doc.Grupo.FaxCodigoPais,
					FaxNumTelefono:     doc.Grupo.FaxNumTelefono,
					CorreoElectronico:  doc.Grupo.CorreoElectronico,
				},
			}
			// crea el emisor en la base de datos
			gErr = tx.Create(&emisor).Error
			if gErr != nil {
				return gErr
			}
			// y actualiza el documento en la base de datos para
			// que apunte a su respectivo Emisor
			gErr = tx.Exec("UPDATE documentos SET emisor_id = ? WHERE documentos.id = ?", emisor.ID, doc.ID).Error
			if gErr != nil {
				return gErr
			}
		}
		return nil
	}
)

func Migrar_15_Crea_Actividades(tx *gorm.DB, m gorm.Migrator) error {
	var a ActividadesEconomicas
	f, err := os.Open("./actividades.json")
	if err != nil {
		return err
	}
	defer f.Close()
	reader := json.NewDecoder(f)
	err = reader.Decode(&a)
	if err != nil {
		return err
	}

	for _, value := range a.Actividades {
		err = tx.Create(&value).Error
		if err != nil {
			return err
		}
	}
	return nil
}
