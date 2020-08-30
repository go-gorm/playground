package v7

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Migrate_20_crear_permisos_y_asocacion_con_usuarios(tx *gorm.DB, m gorm.Migrator) error {
	err := tx.AutoMigrate(&Permiso{}, &Usuario{})
	if err != nil {
		return err
	}
	permisos := []Permiso{
		{Sujeto: "invoice", Accion: "create"},
		{Sujeto: "invoice", Accion: "view"},
		{Sujeto: "admin", Accion: "create"},
		{Sujeto: "admin", Accion: "read"},
		{Sujeto: "admin", Accion: "update"},
		{Sujeto: "admin", Accion: "change"},
		{Sujeto: "admin", Accion: "view"},
		{Sujeto: "issuer", Accion: "change"},
		{Sujeto: "mh_credentials", Accion: "change"},
		{Sujeto: "mh_credentials", Accion: "read"},
		{Sujeto: "user_account", Accion: "change"},
		{Sujeto: "user_account", Accion: "read"},
	}
	for _, permiso := range permisos {
		err = tx.Create(&permiso).Error
		if err != nil {
			return err
		}
	}
	// Mueve los permisos de los grupos a los usuarios
	var suscripciones []Suscripcion
	err = tx.Preload("Usuarios").Find(&suscripciones).Error
	if err != nil {
		return err
	}
	for _, suscripcion := range suscripciones {
		for _, permiso := range suscripcion.Permisos {
			for _, usuario := range suscripcion.Usuarios {
				switch permiso {
				case "admin":
					var permisos_admin []Permiso
					err = tx.Where("sujeto in (?)", []string{"invoice", "admin", "issuer", "mh_credentials", "user_account"}).Find(&permisos_admin).Error
					if err != nil {
						return err
					}
					if len(permisos_admin) == 0 {
						return fmt.Errorf("query a la base de datos no retorno arreglo de permisos esperado para »admin«")
					}
					usuario.Permisos = permisos_admin
					err = tx.Save(&usuario).Error
					if err != nil {
						return err
					}
				case "user":
					err = tx.Where("ID = ?", usuario.ID).Find(&usuario).Error
					if err != nil {
						return err
					}
					if len(usuario.Permisos) == 0 {
						var permisos_user []Permiso
						err = tx.Where("sujeto in (?)", []string{"invoice", "mh_credentials", "user_account"}).Find(&permisos_user).Error
						if err != nil {
							return err
						}
						if len(permisos_user) == 0 {
							return fmt.Errorf("query a la base de datos no retorno arreglo de permisos esperado para »user«")
						}
						usuario.Permisos = permisos_user
						err = tx.Save(&usuario).Error
						if err != nil {
							return err
						}
					}
				default:
					logrus.Warnf("no sé qué hacer con permiso %s", permiso)
				}
			}
		}
	}
	// Elimina el campo permisos de suscripciones
	err = m.DropColumn(&Suscripcion{}, "permisos")
	return err
}
