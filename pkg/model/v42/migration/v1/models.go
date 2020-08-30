package v1migration

import (
	"time"

	"github.com/lib/pq"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/playground/pkg/model"
)

// Emisor representa al emisor del comprobante electrónico
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

// Documento es una estructura que representa los diferentes tipos de
// documentos electrónicos requeridos por Hacienda: Factura electronica, nota
// de débito, nota de crédito, tiquete electrónico. En estos distintos
// documentos no todos los campos son requeridos, algunos siendo opcionales. El
// campo `Tipo` determina que tipo de documento es con el que se esta lidiando.
type Documento struct {
	gorm.Model
	Tipo                           model.DocumentoTipo
	EmisorID                       uint
	ReceptorID                     uint
	Clave                          string                  `xml:"Clave" gorm:"not null;unique;size:50" validate:"max=50,numeric"`
	NumeroConsecutivo              string                  `xml:"NumeroConsecutivo" gorm:"not null;size:20" validate:"max=20,numeric"`
	FechaEmision                   time.Time               `xml:"FechaEmision" gorm:"not null"`
	Emisor                         Emisor                  `xml:"Emisor" gorm:"not null"`
	Receptor                       Receptor                `xml:"Receptor,omitempty"`
	CondicionVenta                 int                     `xml:"CondicionVenta" gorm:"not null" validate:"oneof=1 2 3 4 5 6 99,required"`
	PlazoCredito                   int                     `xml:"PlazoCredito,omitempty"`
	MedioPago                      pq.Int64Array           `xml:"MedioPago" gorm:"type:int[]" validate:"oneof=1 2 3 4 5 99,required"`
	LineaDetalle                   []LineaDetalle          `xml:"LineaDetalle" gorm:"not null"`
	ResumenCodigoMoneda            string                  `xml:"CodigoMoneda,omitempty"` // resumen de factura
	ResumenTipoCambio              decimal.Decimal         `xml:"TipoCambio,omitempty" sql:"type:decimal(18,5);"`
	ResumenTotalServGravados       decimal.Decimal         `xml:"TotalServGravados,omitempty" sql:"type:decimal(18,5);"`
	ResumenTotalServExentos        decimal.Decimal         `xml:"TotalServExentos,omitempty" sql:"type:decimal(18,5);"`
	ResumenTotalMercanciasGravadas decimal.Decimal         `xml:"TotalMercanciasGravadas,omitempty" sql:"type:decimal(18,5);"`
	ResumenTotalMercanciasExentas  decimal.Decimal         `xml:"TotalMercanciasExentas,omitempty" sql:"type:decimal(18,5);"`
	ResumenTotalGravado            decimal.Decimal         `xml:"TotalGravado,omitempty" sql:"type:decimal(18,5);"`
	ResumenTotalExento             decimal.Decimal         `xml:"TotalExento,omitempty" sql:"type:decimal(18,5);"`
	ResumenTotalVenta              decimal.Decimal         `xml:"TotalVenta" gorm:"not null" sql:"type:decimal(18,5);"`
	ResumenTotalDescuentos         decimal.Decimal         `xml:"TotalDescuentos,omitempty" sql:"type:decimal(18,5);"`
	ResumenTotalVentaNeta          decimal.Decimal         `xml:"TotalVentaNeta" gorm:"not null" sql:"type:decimal(18,5);"`
	ResumenTotalImpuesto           decimal.Decimal         `xml:"TotalImpuesto,omitempty" sql:"type:decimal(18,5);"`
	ResumenTotalComprobante        decimal.Decimal         `xml:"TotalComprobante" gorm:"not null" sql:"type:decimal(18,5);"`
	InformacionReferencia          []InformacionReferencia `xml:"InformacionReferencia,omitempty"`
	NormativaNumeroResolucion      string                  `xml:"NumeroResolucion" gorm:"not null"`
	NormativaFechaResolucion       time.Time               `xml:"FechaResolucion" gorm:"not null"`
	SuscripcionID                  uint
}

// Impuesto representa los impuestos en la factura
type Impuesto struct {
	gorm.Model
	Codigo            string          `xml:"Codigo" gorm:"not null" validate:"required"`
	Tarifa            decimal.Decimal `xml:"Tarifa" gorm:"not null" validate:"required" sql:"type:decimal(18,5);"`
	Monto             decimal.Decimal `xml:"Monto" gorm:"not null" validate:"required" sql:"type:decimal(18,5);"`
	TipoDocumento     int             `xml:"TipoDocumento"` // Exoneracion
	NumeroDocumento   string          `xml:"NumeroDocumento"`
	NombreInstitucion string          `xml:"NombreInstitucion"`
	FechaEmision      time.Time       `xml:"FechaEmision"`
	MontoImpuesto     decimal.Decimal `xml:"MontoImpuesto" sql:"type:decimal(18,5);"`
	PorcentajeCompra  int             `xml:"PorcentajeCompra"`
	LineaDetalleID    uint
}

// InformacionReferencia representa una referencia a otros comprobantes electronicos
type InformacionReferencia struct {
	gorm.Model
	TipoDoc      int       `xml:"TipoDoc"`
	Numero       string    `xml:"Numero"`
	FechaEmision time.Time `xml:"FechaEmision"`
	Codigo       int       `xml:"Codigo"`
	Razon        string    `xml:"Razon"`
	DocumentoID  uint
}

// LineaDetalle representa los productos y servicios que se agregan a la factura
type LineaDetalle struct {
	gorm.Model
	NumeroLinea           int             `xml:"NumeroLinea" gorm:"not null" validate:"required"`
	CodigoTipo            int             `xml:"Tipo"`
	CodigoCodigo          string          `xml:"Codigo"`
	Cantidad              decimal.Decimal `xml:"Cantidad" gorm:"not null" validate:"required" sql:"type:decimal(16,3);"`
	UnidadMedida          string          `xml:"UnidadMedida" gorm:"not null" validate:"required"`
	UnidadMedidaComercial string          `xml:"UnidadMedidaComercial,omitempty"`
	Detalle               string          `xml:"Detalle" gorm:"not null" validate:"required"`
	PrecioUnitario        decimal.Decimal `xml:"PrecioUnitario" gorm:"not null" validate:"required" sql:"type:decimal(18,5);"`
	MontoTotal            decimal.Decimal `xml:"MontoTotal" gorm:"not null" validate:"required" sql:"type:decimal(18,5);"`
	MontoDescuento        decimal.Decimal `xml:"MontoDescuento,omitempty" sql:"type:decimal(18,5);"`
	NaturalezaDescuento   string          `xml:"NaturalezaDescuento,omitempty"`
	SubTotal              decimal.Decimal `xml:"SubTotal" gorm:"not null" validate:"required" sql:"type:decimal(18,5);"`
	Impuesto              []Impuesto      `xml:"Impuesto,omitempty"`
	MontoTotalLinea       decimal.Decimal `xml:"MontoTotalLinea" gorm:"not null" validate:"required" sql:"type:decimal(18,5);"`
	DocumentoID           uint
}

type MensajeHacienda struct {
	gorm.Model
	SuscripcionID              uint
	Clave                      string          `xml:"Clave"`
	NombreEmisor               string          `xml:"NombreEmisor"`
	TipoIdentificacionEmisor   int             `xml:"TipoIdentificacionEmisor"`
	NumeroCedulaEmisor         string          `xml:"NumeroCedulaEmisor"`
	NombreReceptor             string          `xml:"NombreReceptor,omitempty"`
	TipoIdentificacionReceptor int             `xml:"TipoIdentificacionReceptor,omitempty"`
	NumeroCedulaReceptor       string          `xml:"NumeroCedulaReceptor,omitempty"`
	Mensaje                    int             `xml:"Mensaje"`
	DetalleMensaje             string          `xml:"DetalleMensaje"`
	MontoTotalImpuesto         decimal.Decimal `xml:"MontoTotalImpuesto,omitempty" sql:"type:decimal(18,5);"`
	TotalFactura               decimal.Decimal `xml:"TotalFactura" sql:"type:decimal(18,5);"`
}

// Receptor representa a un receptor de comprobantes electrónicos, es similar
// al tipo Emisor con la diferencia de que este tipo añade un campo adicional
type Receptor struct {
	gorm.Model
	Nombre                   string `xml:"Nombre"`
	IdentificacionTipo       int    `xml:"Tipo"`
	IdentificacionNumero     string `xml:"Numero" gorm:"unique"`
	NombreComercial          string `xml:"NombreComercial,omitempty"`
	Provincia                int    `xml:"Provincia"`
	Canton                   int    `xml:"Canton"`
	Distrito                 int    `xml:"Distrito"`
	Barrio                   int    `xml:"Barrio,omitempty"`
	OtrasSenas               string `xml:"OtrasSenas"`
	TelCodigoPais            int    `xml:"CodigoPais"`
	TelNumTelefono           string `xml:"NumTelefono"`
	FaxCodigoPais            int
	FaxNumTelefono           string
	CorreoElectronico        string `xml:"CorreoElectronico"`
	IdentificacionExtranjero string
}

// Usuario representa los datos de autenticación de los usuarios de la
// aplicación web
type Usuario struct {
	gorm.Model
	SuscripcionID int
	Cuenta        string `gorm:"not null;unique" validate:"email"`
	Clave         string `gorm:"not null" validate:"min=8"`
}

// OAuthResponse respresenta la respuesta que da el servidor OAuth de Hacienda
type OAuthResponse struct {
	ID               uint      `gorm:"primary_key" json:"-"`
	CreatedAt        time.Time `json:"-"`
	UpdatedAt        time.Time `json:"-"`
	AccessToken      string    `json:"access_token"`
	ExpiresIn        int       `json:"expires_in"`
	RefreshExpiresIn int       `json:"refresh_expires_in"`
	RefreshToken     string    `json:"refresh_token"`
}

// Credencial guarda información de autenticación de un obligado
// contribuyente procedientes de Hacienda
type Credencial struct {
	gorm.Model
	Usuario    string `gorm:"not null;unique"`
	Clave      string `gorm:"not null"`
	Pin        string `gorm:"not null"`
	ArchivoP12 string `gorm:"not null;unique"`
}

// Grupo conserva información para la emisión de facturas
type Grupo struct {
	gorm.Model
	Nombre               string `gorm:"not null;size:80;unique"`
	IdentificacionTipo   int    `gorm:"not null"`
	IdentificacionNumero string `gorm:"not null;size:14;unique"`
	NombreComercial      string `gorm:"size:80" validate:"max=80"`
	Provincia            int    `gorm:"not null"`
	Canton               int    `gorm:"not null"`
	Distrito             int    `gorm:"not null"`
	Barrio               int
	OtrasSenas           string `gorm:"size:180"`
	TelCodigoPais        int
	TelNumTelefono       string `gorm:"size:20"`
	FaxCodigoPais        int
	FaxNumTelefono       string `gorm:"size:20"`
	CorreoElectronico    string `gorm:"not null;size:60"`
	NumeroDocumento      int64
	OAuth                OAuthResponse `gorm:"foreignkey:OAuthResponseID"`
	OAuthResponseID      uint
	Credenciales         Credencial `gorm:"foreignkey:CredencialID"`
	CredencialID         uint
	SuscripcionID        uint
}

// Suscripcion representa una suscripción al servicio de generación de
// comprobantes electrónicos, también contiene permisos de acceso para
// distintas areas de la aplicación web
type Suscripcion struct {
	gorm.Model
	Nombre               string `gorm:"not null;size:80;unique"`
	IdentificacionTipo   int    `gorm:"not null"`
	IdentificacionNumero string `gorm:"not null;size:14;unique"`
	Usuarios             []Usuario
	Grupos               []Grupo
	Documentos           []Documento
	Mensajes             []MensajeHacienda
	Permisos             pq.StringArray        `gorm:"type:varchar(64)[]"`
	Tipo                 model.SuscripcionTipo `gorm:"type:int"`
}
