package v11migration

import (
	"time"

	"github.com/lib/pq"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/playground/pkg/model"
)

// Documento es una estructura que representa los diferentes tipos de
// documentos electrónicos requeridos por Hacienda: Factura electronica, nota
// de débito, nota de crédito, tiquete electrónico. En estos distintos
// documentos no todos los campos son requeridos, algunos siendo opcionales. El
// campo `Tipo` determina que tipo de documento es con el que se esta lidiando.
type Documento struct {
	gorm.Model
	Tipo                           model.DocumentoTipo     `json:"tipo" validate:"oneof=1 2 3 4,required"`
	Intento                        int                     `json:"intento"`
	Caja                           int                     `json:"caja" gorm:"-"`
	Sucursal                       int                     `json:"sucursal" gorm:"-"`
	Estado                         model.EstadoTipo        `json:"estado"`
	GrupoID                        uint                    `json:"grupo_id"`
	ReceptorID                     uint                    `json:"receptor_id"`
	Locacion                       string                  `json:"locacion"`
	URLXML                         string                  `json:"urlxml" gorm:"-"`
	URLXMLConfirmacion             string                  `json:"urlxml_confirmacion" gorm:"-"`
	Clave                          string                  `json:"clave" gorm:"not null;unique;size:50"`
	NumeroConsecutivo              string                  `json:"numero_consecutivo" gorm:"not null;size:20"`
	FechaEmision                   model.CustomTime        `json:"fecha_emision" gorm:"not null"`
	Grupo                          Grupo                   `json:"grupo" gorm:"not null" validate:"dive,required" xml:"Emisor"`
	Receptor                       Receptor                `json:"receptor"`
	CondicionVenta                 int                     `gorm:"not null" validate:"oneof=1 2 3 4 5 6 99,required" json:"condicion_venta"`
	PlazoCredito                   int                     `json:"plazo_credito"`
	MedioPago                      pq.Int64Array           `gorm:"type:int[]" validate:"dive,oneof=1 2 3 4 5 99,required" json:"medio_pago"`
	LineaDetalle                   []LineaDetalle          `gorm:"not null,PRELOAD:false" json:"linea_detalle" validate:"required,min=1,max=1000" xml:"DetalleServicio>LineaDetalle"`
	ResumenCodigoMoneda            string                  `json:"resumen_codigo_moneda" xml:"ResumenFactura>CodigoMoneda"` // resumen de factura
	ResumenTipoCambio              decimal.Decimal         `sql:"type:decimal(18,5);" json:"resumen_tipo_cambio" xml:"ResumenFactura>TipoCambio"`
	ResumenTotalServGravados       decimal.Decimal         `sql:"type:decimal(18,5);" json:"resumen_total_serv_gravados" xml:"ResumenFactura>TotalServGravados"`
	ResumenTotalServExentos        decimal.Decimal         `sql:"type:decimal(18,5);" json:"resumen_total_serv_exentos" xml:"ResumenFactura>TotalServExcentos"`
	ResumenTotalMercanciasGravadas decimal.Decimal         `sql:"type:decimal(18,5);" json:"resumen_total_mercancias_gravadas" xml:"ResumenFactura>TotalMercanciasGravadas"`
	ResumenTotalMercanciasExentas  decimal.Decimal         `sql:"type:decimal(18,5);" json:"resumen_total_mercancias_exentas" xml:"ResumenFactura>TotalMercanciasExentas"`
	ResumenTotalGravado            decimal.Decimal         `sql:"type:decimal(18,5);" json:"resumen_total_gravado" xml:"ResumenFactura>TotalGravado"`
	ResumenTotalExento             decimal.Decimal         `sql:"type:decimal(18,5);" json:"resumen_total_exento" xml:"ResumenFactura>TotalExento"`
	ResumenTotalVenta              decimal.Decimal         `gorm:"not null" sql:"type:decimal(18,5);" json:"resumen_total_venta" xml:"ResumenFactura>TotalVenta"`
	ResumenTotalDescuentos         decimal.Decimal         `sql:"type:decimal(18,5);" json:"resumen_total_descuentos" xml:"ResumenFactura>TotalDescuentos"`
	ResumenTotalVentaNeta          decimal.Decimal         `gorm:"not null" sql:"type:decimal(18,5);" json:"resumen_total_venta_neta" xml:"ResumenFactura>TotalVentaNeta"`
	ResumenTotalImpuesto           decimal.Decimal         `sql:"type:decimal(18,5);" json:"resumen_total_impuesto" xml:"ResumenFactura>TotalImpuesto"`
	ResumenTotalComprobante        decimal.Decimal         `gorm:"not null" sql:"type:decimal(18,5);" json:"resumen_total_comprobante" xml:"ResumenFactura>TotalComprobante"`
	InformacionReferencia          []InformacionReferencia `json:"informacion_referencia" xml:"InformacionReferencia"`
	NormativaNumeroResolucion      string                  `gorm:"not null" json:"normativa_numero_resolucion" xml:"Normativa>NumeroResolucion"`
	NormativaFechaResolucion       model.CustomTime        `gorm:"not null" json:"normativa_fecha_resolucion" xml:"Normativa>FechaResolucion"`
}

// Grupo conserva información para la emisión de facturas
type Grupo struct {
	gorm.Model
	Nombre               string            `gorm:"not null;size:80;unique" json:"nombre" validate:"required"`
	IdentificacionTipo   int               `gorm:"not null" json:"identificacion_tipo" xml:"Identificacion>Tipo"`
	IdentificacionNumero string            `gorm:"not null;size:14;unique" json:"identificacion_numero" validate:"required" xml:"Identificacion>Numero"`
	NombreComercial      string            `gorm:"size:80" validate:"max=80" json:"nombre_comercial"`
	Provincia            int               `gorm:"not null" json:"provincia" xml:"Ubicacion>Provincia"`
	Canton               int               `gorm:"not null" json:"canton" xml:"Ubicacion>Canton"`
	Distrito             int               `gorm:"not null" json:"distrito" xml:"Ubicacion>Distrito"`
	Barrio               int               `json:"barrio" xml:"Ubicacion>Barrio"`
	OtrasSenas           string            `gorm:"size:180" json:"otras_senas" xml:"Ubicacion>OtrasSenas"`
	TelCodigoPais        int               `json:"tel_codigo_pais" xml:"Telefono>CodigoPais"`
	TelNumTelefono       string            `gorm:"size:20" json:"tel_num_telefono" xml:"Telefono>NumTelefono"`
	FaxCodigoPais        int               `json:"fax_codigo_pais" xml:"Fax>CodigoPais"`
	FaxNumTelefono       string            `gorm:"size:20" json:"fax_num_telefono" xml:"Fax>NumTelefono"`
	CorreoElectronico    string            `gorm:"not null;size:60" json:"correo_electronico" validate:"required,email"`
	OAuth                OAuthResponse     `gorm:"foreignkey:OAuthResponseID" json:"o_auth"`
	Credenciales         Credencial        `gorm:"foreignkey:CredencialID" json:"credenciales"`
	CodigoContent        string            `gorm:"-" json:"codigo_content"`
	Codigos              []Codigo          `json:"codigos"`
	OAuthResponseID      uint              `json:"o_auth_response_id"`
	CredencialID         uint              `json:"credencial_id"`
	SuscripcionID        uint              `json:"suscripcion_id"`
	Documentos           []Documento       `json:"documentos"`
	Mensajes             []MensajeHacienda `json:"mensajes"`
}

// Receptor representa a un receptor de comprobantes electrónicos, es similar
// al tipo Emisor con la diferencia de que este tipo añade un campo adicional
type Receptor struct {
	gorm.Model
	Nombre                   string `json:"nombre"`
	IdentificacionTipo       int    `json:"identificacion_tipo" xml:"Identificacion>Tipo"`
	IdentificacionNumero     string `gorm:"unique" json:"identificacion_numero" xml:"Identificacion>Numero"`
	NombreComercial          string `json:"nombre_comercial"`
	Provincia                int    `json:"provincia" xml:"Ubicacion>Provincia"`
	Canton                   int    `json:"canton" xml:"Ubicacion>Canton"`
	Distrito                 int    `json:"distrito" xml:"Ubicacion>Distrito"`
	Barrio                   int    `json:"barrio" xml:"Ubicacion>Barrio"`
	OtrasSenas               string `json:"otras_senas" xml:"Ubicacion>OtrasSenas"`
	TelCodigoPais            int    `json:"tel_codigo_pais" xml:"Telefono>CodigoPais"`
	TelNumTelefono           string `json:"tel_num_telefono" xml:"Telefono>NumTelefono"`
	FaxCodigoPais            int    `json:"fax_codigo_pais" xml:"Fax>CodigoPais"`
	FaxNumTelefono           string `json:"fax_num_telefono" xml:"Fax>NumTelefono"`
	CorreoElectronico        string `json:"correo_electronico"`
	IdentificacionExtranjero string `json:"identificacion_extranjero"`
}

// InformacionReferencia representa una referencia a otros comprobantes electronicos
type InformacionReferencia struct {
	gorm.Model
	TipoDoc      int              `json:"tipo_doc" xml:"TipoDoc"`
	Numero       string           `json:"numero" xml:"Numero"`
	FechaEmision model.CustomTime `json:"fecha_emision" xml:"FechaEmision"`
	Codigo       int              `json:"codigo" xml:"Codigo"`
	Razon        string           `json:"razon" xml:"Razon"`
	DocumentoID  uint             `json:"documento_id"`
}

type Recepcion struct {
	Documento
	EstadoRecepcion model.EstadoReceptorTipo
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

// Codigo es una tabla que enlaza distintos códigos con detalles
type Codigo struct {
	gorm.Model
	Codigo  string `gorm:"not null;size:80" json:"codigo"`
	Detalle string `gorm:"not null;size:160" json:"detalle"`
	GrupoID uint   `json:"grupo_id"`
}

// MensajeHacienda es una estructura en XML que Hacienda emite, según la
// documentación: Establece los archivos XML para los mensajes que deben de
// utilizar los obligados tributarios al momento de la confirmación de
// aceptación o rechazo de los documentos electrónicos, así como el mensaje que
// utilizará la Dirección General de Tributacion para comunicar al obligado
// tributarios la validación del comprobante electrónico. Estos mensajes deben
// estar firmados digitalmente por el receptor del comprobante electrónico o
// por el Ministerio de Hacienda, según corresponda por el tipo de mensaje.
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

// LineaDetalle representa los productos y servicios que se agregan a la factura
type LineaDetalle struct {
	gorm.Model
	NumeroLinea           int             `gorm:"not null" validate:"required,numeric,min=1" json:"numero_linea"`
	CodigoTipo            int             `json:"codigo_tipo" validate:"oneof=1 2 3 4 99" xml:"Codigo>Tipo"`
	Codigo                string          `json:"codigo" validate:"required,max=20" xml:"Codigo>Codigo"`
	Cantidad              decimal.Decimal `gorm:"not null" sql:"type:decimal(16,3);" json:"cantidad" validate:"required,numeric,max=9999999999999999.999"`
	UnidadMedida          string          `gorm:"not null" validate:"required" json:"unidad_medida"`
	UnidadMedidaComercial string          `json:"unidad_medida_comercial"`
	Detalle               string          `gorm:"not null" validate:"required,max=160" json:"detalle"`
	PrecioUnitario        decimal.Decimal `gorm:"not null" validate:"required,numeric,max=9999999999999.99999" sql:"type:decimal(18,5);" json:"precio_unitario"`
	MontoTotal            decimal.Decimal `gorm:"not null" validate:"required,numeric,max=9999999999999.99999" sql:"type:decimal(18,5);" json:"monto_total"`
	MontoDescuento        decimal.Decimal `sql:"type:decimal(18,5);" json:"monto_descuento" validate:"omitempty,numeric,max=9999999999999.99999"`
	NaturalezaDescuento   string          `json:"naturaleza_descuento" validate:"max=80"`
	SubTotal              decimal.Decimal `gorm:"not null" validate:"required,numeric,max=9999999999999.99999" sql:"type:decimal(18,5);" json:"sub_total"`
	Impuesto              []Impuesto      `json:"impuesto" validate:"omitempty"`
	MontoTotalLinea       decimal.Decimal `gorm:"not null" validate:"required,numeric,max=9999999999999.99999" sql:"type:decimal(18,5);" json:"monto_total_linea"`
	DocumentoID           uint            `json:"documento_id"`
}

// Impuesto representa los impuestos en la factura
type Impuesto struct {
	gorm.Model
	Codigo            int             `gorm:"not null" validate:"required" json:"codigo"`
	Tarifa            decimal.Decimal `gorm:"not null" validate:"required" sql:"type:decimal(18,5);" json:"tarifa"`
	Monto             decimal.Decimal `gorm:"not null" validate:"required" sql:"type:decimal(18,5);" json:"monto"`
	TipoDocumento     int             `json:"tipo_documento" xml:"Exoneracion>TipoDocumento"` // Exoneracion
	NumeroDocumento   string          `json:"numero_documento" xml:"Exoneracion>NumeroDocumento"`
	NombreInstitucion string          `json:"nombre_institucion" xml:"Exoneracion>NombreInstitucion"`
	FechaEmision      time.Time       `json:"fecha_emision" xml:"Exoneracion>FechaEmision"`
	MontoImpuesto     decimal.Decimal `sql:"type:decimal(18,5);" json:"monto_impuesto" xml:"Exoneracion>MontoImpuesto"`
	PorcentajeCompra  int             `json:"porcentaje_compra" xml:"Exoneracion>PorcentajeCompra"`
	LineaDetalleID    uint            `json:"linea_detalle_id"`
}
