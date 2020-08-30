package v3

import (
	"encoding/json"
	"os"

	"gorm.io/gorm"
)

func Migrate_16_ubicaciones(tx *gorm.DB, m gorm.Migrator) error {
	// estructura para extraer datos del archivo JSON
	type Ubicaciones struct {
		Cantones  []Canton   `json:"cantones"`
		Distritos []Distrito `json:"distritos"`
		Barrios   []Barrio   `json:"barrios"`
	}
	f, err := os.Open("./ubicaciones.json")
	if err != nil {
		return err
	}
	defer f.Close()
	reader := json.NewDecoder(f)
	var u Ubicaciones
	err = reader.Decode(&u)
	if err != nil {
		return err
	}
	// crea las tablas
	err = tx.AutoMigrate(&Canton{}, &Distrito{}, &Barrio{})
	if err != nil {
		return err
	}
	for _, value := range u.Cantones {
		dbErr := tx.Create(&value).Error
		if dbErr != nil {
			return dbErr
		}
	}
	for _, value := range u.Distritos {
		dbErr := tx.Create(&value).Error
		if dbErr != nil {
			return dbErr
		}
	}
	for _, value := range u.Barrios {
		dbErr := tx.Create(&value).Error
		if dbErr != nil {
			return dbErr
		}
	}
	return nil
}

func Rollback_16_ubicaciones(tx *gorm.DB, m gorm.Migrator) error {
	err := m.DropTable(&Canton{}, &Distrito{}, &Barrio{})
	return err
}
