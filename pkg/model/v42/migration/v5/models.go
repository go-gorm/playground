package v5migration

import "gorm.io/playground/pkg/model"

type Documento struct {
	Estado   model.EstadoTipo
	Locacion string
}
