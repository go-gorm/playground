package model42

import (
	"time"

	"github.com/lib/pq"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/playground/pkg/model"
)

type Documento struct {
	gorm.Model
	Tipo                           model.DocumentoTipo
	Intento                        int
	Caja                           int `gorm:"-"`
	Sucursal                       int `gorm:"-"`
	Estado                         model.EstadoTipo
	GrupoID                        uint
	ReceptorID                     uint
	Locacion                       string
	URLXML                         string    `gorm:"-"`
	URLXMLConfirmacion             string    `gorm:"-"`
	Clave                          string    `gorm:"not null;unique;size:50"`
	NumeroConsecutivo              string    `gorm:"not null;size:20"`
	FechaEmision                   time.Time `gorm:"not null"`
	Grupo                          Grupo     `gorm:"not null"`
	Receptor                       Receptor
	CondicionVenta                 int `gorm:"not null"`
	PlazoCredito                   int
	MedioPago                      pq.Int64Array   `gorm:"type:int[]"`
	LineaDetalle                   []LineaDetalle  `gorm:"not null,PRELOAD:false"`
	ResumenCodigoMoneda            string          // resumen de factura
	ResumenTipoCambio              decimal.Decimal `sql:"type:decimal(18,5);"`
	ResumenTotalServGravados       decimal.Decimal `sql:"type:decimal(18,5);"`
	ResumenTotalServExentos        decimal.Decimal `sql:"type:decimal(18,5);"`
	ResumenTotalMercanciasGravadas decimal.Decimal `sql:"type:decimal(18,5);"`
	ResumenTotalMercanciasExentas  decimal.Decimal `sql:"type:decimal(18,5);"`
	ResumenTotalGravado            decimal.Decimal `sql:"type:decimal(18,5);"`
	ResumenTotalExento             decimal.Decimal `sql:"type:decimal(18,5);"`
	ResumenTotalVenta              decimal.Decimal `gorm:"not null" sql:"type:decimal(18,5);"`
	ResumenTotalDescuentos         decimal.Decimal `sql:"type:decimal(18,5);"`
	ResumenTotalVentaNeta          decimal.Decimal `gorm:"not null" sql:"type:decimal(18,5);"`
	ResumenTotalImpuesto           decimal.Decimal `sql:"type:decimal(18,5);"`
	ResumenTotalComprobante        decimal.Decimal `gorm:"not null" sql:"type:decimal(18,5);"`
	InformacionReferencia          []InformacionReferencia
	NormativaNumeroResolucion      string    `gorm:"not null"`
	NormativaFechaResolucion       time.Time `gorm:"not null"`
}

type Impuesto struct {
	gorm.Model
	Codigo            int             `gorm:"not null"`
	Tarifa            decimal.Decimal `gorm:"not null" sql:"type:decimal(18,5);"`
	Monto             decimal.Decimal `gorm:"not null" sql:"type:decimal(18,5);"`
	TipoDocumento     int             // Exoneracion
	NumeroDocumento   string
	NombreInstitucion string
	FechaEmision      time.Time
	MontoImpuesto     decimal.Decimal `sql:"type:decimal(18,5);"`
	PorcentajeCompra  int
	LineaDetalleID    uint
}

type InformacionReferencia struct {
	gorm.Model
	TipoDoc      int
	Numero       string
	FechaEmision time.Time
	Codigo       int
	Razon        string
	DocumentoID  uint
}

// LineaDetalle representa los productos y servicios que se agregan a la factura
type LineaDetalle struct {
	gorm.Model
	NumeroLinea           int `gorm:"not null"`
	CodigoTipo            int
	Codigo                string
	Cantidad              decimal.Decimal `gorm:"not null" sql:"type:decimal(16,3);"`
	UnidadMedida          string          `gorm:"not null"`
	UnidadMedidaComercial string
	Detalle               string          `gorm:"not null"`
	PrecioUnitario        decimal.Decimal `gorm:"not null" sql:"type:decimal(18,5);"`
	MontoTotal            decimal.Decimal `gorm:"not null" sql:"type:decimal(18,5);"`
	MontoDescuento        decimal.Decimal `sql:"type:decimal(18,5);"`
	NaturalezaDescuento   string
	SubTotal              decimal.Decimal `gorm:"not null" sql:"type:decimal(18,5);"`
	Impuesto              []Impuesto
	MontoTotalLinea       decimal.Decimal `gorm:"not null" sql:"type:decimal(18,5);"`
	DocumentoID           uint
}

type MensajeHacienda struct {
	gorm.Model
	GrupoID                    uint
	Clave                      string
	NumeroConsecutivoReceptor  string
	NombreEmisor               string
	TipoIdentificacionEmisor   int
	NumeroCedulaEmisor         string
	NombreReceptor             string
	TipoIdentificacionReceptor int
	NumeroCedulaReceptor       string
	Mensaje                    int
	DetalleMensaje             string
	MontoTotalImpuesto         decimal.Decimal `sql:"type:decimal(18,5);"`
	TotalFactura               decimal.Decimal `sql:"type:decimal(18,5);"`
}

type Receptor struct {
	gorm.Model
	Nombre                   string
	IdentificacionTipo       int
	IdentificacionNumero     string `gorm:"unique"`
	NombreComercial          string
	Provincia                int
	Canton                   int
	Distrito                 int
	Barrio                   int
	OtrasSenas               string
	TelCodigoPais            int
	TelNumTelefono           string
	FaxCodigoPais            int
	FaxNumTelefono           string
	CorreoElectronico        string
	IdentificacionExtranjero string
}

type Usuario struct {
	gorm.Model
	SuscripcionID int
	Cuenta        string `gorm:"not null;unique"`
	Clave         string `gorm:"not null"`
	Plano         string `gorm:"-"`
}

type OAuthResponse struct {
	ID               uint `gorm:"primary_key"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	AccessToken      string
	ExpiresIn        int
	RefreshExpiresIn int
	RefreshToken     string
}

type Credencial struct {
	gorm.Model
	Usuario    string `gorm:"not null;unique"`
	Clave      string `gorm:"not null"`
	Pin        string `gorm:"not null"`
	ArchivoP12 string `gorm:"not null;unique"`
	Contenido  string `gorm:"-"`
}

type Grupo struct {
	gorm.Model
	Nombre               string `gorm:"not null;size:80;unique"`
	IdentificacionTipo   int    `gorm:"not null"`
	IdentificacionNumero string `gorm:"not null;size:14;unique"`
	NombreComercial      string `gorm:"size:80"`
	Provincia            int    `gorm:"not null"`
	Canton               int    `gorm:"not null"`
	Distrito             int    `gorm:"not null"`
	Barrio               int
	OtrasSenas           string `gorm:"size:180"`
	TelCodigoPais        int
	TelNumTelefono       string `gorm:"size:20"`
	FaxCodigoPais        int
	FaxNumTelefono       string        `gorm:"size:20"`
	CorreoElectronico    string        `gorm:"not null;size:60"`
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

type Suscripcion struct {
	gorm.Model
	Nombre               string `gorm:"not null;size:80;unique"`
	IdentificacionTipo   int    `gorm:"not null"`
	IdentificacionNumero string `gorm:"not null;size:14;unique"`
	Usuarios             []Usuario
	Grupos               []Grupo
	Permisos             pq.StringArray        `gorm:"type:varchar(64)[]"`
	Tipo                 model.SuscripcionTipo `gorm:"type:int"`
}

type Codigo struct {
	gorm.Model
	Codigo  string `gorm:"not null;size:80;uniqueIndex:idx_codigo_grupoid"`
	Detalle string `gorm:"not null;size:160"`
	GrupoID uint   `gorm:"uniqueIndex:idx_codigo_grupoid"`
}

type Recepcion struct {
	Documento
	EstadoRecepcion model.EstadoReceptorTipo
}

type Consecutivo struct {
	gorm.Model
	GrupoID  uint                  `gorm:"uniqueIndex:idx_grupo_tipo_caja_sucursal"`
	Tipo     model.ConsecutivoTipo `gorm:"uniqueIndex:idx_grupo_tipo_caja_sucursal"`
	Sucursal int                   `gorm:"uniqueIndex:idx_grupo_tipo_caja_sucursal"`
	Caja     int                   `gorm:"uniqueIndex:idx_grupo_tipo_caja_sucursal"`
	Contador int64                 `gorm:"default:'1'"`
}
