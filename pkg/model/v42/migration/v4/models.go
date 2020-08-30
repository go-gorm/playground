package v4migration

import "gorm.io/gorm"

type Emisor struct {
	gorm.Model
	Nombre               string `xml:"Nombre" gorm:"not null;size:80" validate:"max=80"`
	IdentificacionTipo   int    `xml:"Tipo" gorm:"not null"`
	IdentificacionNumero string `xml:"Numero" gorm:"not null;size:14" validate:"max=14"`
	NombreComercial      string `xml:"NombreComercial,omitempty" gorm:"size:80" validate:"max=80"`
	Provincia            int    `xml:"Provincia" gorm:"not null"`
	Canton               int    `xml:"Canton" gorm:"not null"`
	Distrito             int    `xml:"Distrito" gorm:"not null"`
	Barrio               int    `xml:"Barrio,omitempty"`
	OtrasSenas           string `xml:"OtrasSenas" gorm:"size:180" validate:"max=180"`
	TelCodigoPais        int    `xml:"CodigoPais"`
	TelNumTelefono       string `xml:"NumTelefono" gorm:"size:20" validate:"max=20"`
	FaxCodigoPais        int
	FaxNumTelefono       string `validate:"max=20" gorm:"size:20"`
	CorreoElectronico    string `xml:"CorreoElectronico" gorm:"not null;size:60" validate:"email;size:60"`
}
type Documento struct {
	EmisorID uint
	GrupoID  uint
}

type Grupo struct {
	Documentos []Documento
	Mensajes   []MensajeHacienda
}

type MensajeHacienda struct {
	GrupoID uint
}
