package v2

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Actividad struct {
	gorm.Model
	Codigo      string   `gorm:"unique;not null;size=6"`
	Descripcion string   `gorm:"not null"`
	Grupos      []*Grupo `gorm:"many2many:ga"`
}

type GrupoBase struct {
	IdentificacionTipo int    `gorm:"not null"`
	NombreComercial    string `gorm:"size:80"`
	Provincia          int    `gorm:"not null"`
	Canton             int    `gorm:"not null"`
	Distrito           int    `gorm:"not null"`
	Barrio             int
	OtrasSenas         string `gorm:"size:250"`
	TelCodigoPais      int
	TelNumTelefono     string `gorm:"size:20"`
	FaxCodigoPais      int
	FaxNumTelefono     string `gorm:"size:20"`
	CorreoElectronico  string `gorm:"not null;size:60"`
}

// OAuthResponse respresenta la respuesta que da el servidor OAuth de Hacienda
type OAuthResponse struct {
	ID               uint      `gorm:"primary_key" json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	AccessToken      string    `json:"access_token"`
	ExpiresIn        int       `json:"expires_in"`
	RefreshExpiresIn int       `json:"refresh_expires_in"`
	RefreshToken     string    `json:"refresh_token"`
}

type Credencial struct {
	gorm.Model
	Usuario    string `gorm:"not null;unique" json:"usuario"`
	Clave      string `gorm:"not null" json:"clave"`
	Pin        string `gorm:"not null" json:"pin"`
	ArchivoP12 string `gorm:"not null;unique" json:"archivo_p_12"`
	Contenido  string `gorm:"-" json:"contenido"`
}

type Codigo struct {
	gorm.Model
	Codigo  string `gorm:"not null;size:80" json:"codigo"`
	Detalle string `gorm:"not null;size:160" json:"detalle"`
	GrupoID uint   `json:"grupo_id"`
}

type Grupo struct {
	gorm.Model
	GrupoBase
	Nombre               string        `gorm:"not null;size:100;unique"`
	IdentificacionNumero string        `gorm:"not null;size:14;unique"`
	Actividades          []*Actividad  `gorm:"many2many:ga"`
	OAuth                OAuthResponse `gorm:"foreignkey:OAuthResponseID"`
	Credenciales         Credencial    `gorm:"foreignkey:CredencialID"`
	CodigoContent        string        `gorm:"-"`
	Codigos              []Codigo
	OAuthResponseID      uint
	CredencialID         uint
	SuscripcionID        uint
	Documentos           []Documento
	Mensajes             []MensajeHacienda
}

type Emisor struct {
	gorm.Model
	GrupoBase
	Nombre               string `gorm:"not null;size:100"`
	IdentificacionNumero string `gorm:"not null;size:14"`
}

type Documento struct {
	ID          uint
	EmisorID    uint
	Emisor      Emisor
	ActividadID uint
	Actividad   Actividad
}

type MensajeHacienda struct {
	gorm.Model
	GrupoID                    uint            `json:"grupo_id"`
	Clave                      string          `json:"clave"`
	NumeroConsecutivoReceptor  string          `json:"numero_consecutivo_receptor"`
	NombreEmisor               string          `json:"nombre_emisor"`
	TipoIdentificacionEmisor   int             `json:"tipo_identificacion_emisor"`
	NumeroCedulaEmisor         string          `json:"numero_cedula_emisor"`
	NombreReceptor             string          `json:"nombre_receptor"`
	TipoIdentificacionReceptor int             `json:"tipo_identificacion_receptor"`
	NumeroCedulaReceptor       string          `json:"numero_cedula_receptor"`
	Mensaje                    int             `json:"mensaje"`
	DetalleMensaje             string          `json:"detalle_mensaje"`
	MontoTotalImpuesto         decimal.Decimal `sql:"type:decimal(18,5);" json:"monto_total_impuesto"`
	TotalFactura               decimal.Decimal `sql:"type:decimal(18,5);" json:"total_factura"`
}
