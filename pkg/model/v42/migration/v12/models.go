package v12migration

import (
	"gorm.io/gorm"
	"gorm.io/playground/pkg/model"
)

type Grupo struct {
	gorm.Model
	NumeroDocumento   int64
	NumeroDocumentoNC int64
	NumeroDocumentoND int64
	NumeroDocumentoTC int64
}

type Consecutivo struct {
	gorm.Model
	GrupoID  uint                  `json:"grupo_id" gorm:"uniqueIndex:idx_grupo_tipo_caja_sucursal"`
	Tipo     model.ConsecutivoTipo `json:"tipo" gorm:"uniqueIndex:idx_grupo_tipo_caja_sucursal"`
	Sucursal int                   `json:"sucursal" gorm:"uniqueIndex:idx_grupo_tipo_caja_sucursal"`
	Caja     int                   `json:"caja" gorm:"uniqueIndex:idx_grupo_tipo_caja_sucursal"`
	Contador int64                 `json:"contador"`
}
