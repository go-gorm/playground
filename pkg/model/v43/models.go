package models43

import (
	"time"

	"gitlab.com/shackra/gormbug/pkg/model"

	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Documento es una estructura que representa los diferentes tipos de
// documentos electrónicos requeridos por Hacienda: Factura electronica, nota
// de débito, nota de crédito, tiquete electrónico. En estos distintos
// documentos no todos los campos son requeridos, algunos siendo opcionales. El
// campo `Tipo` determina que tipo de documento es con el que se esta lidiando.
type Documento struct {
	gorm.Model
	Tipo                           model.DocumentoTipo     `ts:"ComprobanteTipo" json:"tipo" validate:"oneof=1 2 3 4 5 6,required"`
	Intento                        int                     `ts:",null" json:"intento"`
	Caja                           int                     `json:"caja" gorm:"-"`
	Sucursal                       int                     `json:"sucursal" gorm:"-"`
	Estado                         model.EstadoTipo        `ts:"ComprobanteEstado,null" json:"estado"`
	EmisorID                       uint                    `ts:"-" json:"emisor_id"`
	Emisor                         Emisor                  `json:"emisor" gorm:"not null" validate:"dive,required"`
	ActividadID                    uint                    `ts:"-" json:"actividad_id"`
	Actividad                      Actividad               `json:"actividad"`
	ReceptorID                     uint                    `ts:"-" json:"receptor_id"`
	Locacion                       string                  `json:"locacion"`
	URLXML                         string                  `ts:",null" json:"urlxml" gorm:"-"`
	URLXMLConfirmacion             string                  `ts:",null" json:"urlxml_confirmacion" gorm:"-"`
	Clave                          string                  `ts:",null" json:"clave" gorm:"not null;unique;size:50"`
	NumeroConsecutivo              string                  `ts:",null" json:"numero_consecutivo" gorm:"not null;size:20"`
	FechaEmision                   model.CustomTime        `ts:"date,null" json:"fecha_emision" gorm:"not null"`
	Grupo                          Grupo                   `ts:"-" json:"-" gorm:"not null" xml:"-" validate:"-"`
	GrupoID                        uint                    `ts:"-" json:"grupo_id"`
	Receptor                       Receptor                `ts:",null" json:"receptor" validate:"dive,omitempty"`
	CondicionVenta                 int                     `ts:"CondicionVentaTipo" gorm:"not null" validate:"oneof=1 2 3 4 5 6 99,required" json:"condicion_venta"`
	PlazoCredito                   int                     `json:"plazo_credito"`
	MedioPago                      pq.Int64Array           `ts:"FormasDePago[]" gorm:"type:int[]" validate:"dive,oneof=1 2 3 4 5 99,required" json:"medio_pago"`
	LineaDetalle                   []LineaDetalle          `gorm:"not null,PRELOAD:false" json:"linea_detalle" validate:"required,min=1,max=1000" xml:"DetalleServicio>LineaDetalle"`
	OtrosCargos                    []OtrosCargos           `ts:"null" json:"otros_cargos"`
	ResumenCodigoMoneda            string                  `ts:"Monedas" json:"resumen_codigo_moneda" xml:"ResumenFactura>CodigoMoneda"` // resumen de factura
	ResumenTipoCambio              decimal.Decimal         `ts:"decimal" sql:"type:decimal(18,5);" json:"resumen_tipo_cambio" xml:"ResumenFactura>TipoCambio"`
	ResumenTotalServGravados       decimal.Decimal         `ts:"decimal,null" sql:"type:decimal(18,5);" json:"resumen_total_serv_gravados" xml:"ResumenFactura>TotalServGravados"`
	ResumenTotalServExentos        decimal.Decimal         `ts:"decimal,null" sql:"type:decimal(18,5);" json:"resumen_total_serv_exentos" xml:"ResumenFactura>TotalServExcentos"`
	ResumenTotalServExonerado      decimal.Decimal         `ts:"decimal,null" sql:"type:decimal(18,5);" json:"resumen_total_serv_exonerado" xml:"ResumenFactura>TotalServExonerado"`
	ResumenTotalMercanciasGravadas decimal.Decimal         `ts:"decimal,null" sql:"type:decimal(18,5);" json:"resumen_total_mercancias_gravadas" xml:"ResumenFactura>TotalMercanciasGravadas"`
	ResumenTotalMercanciasExentas  decimal.Decimal         `ts:"decimal,null" sql:"type:decimal(18,5);" json:"resumen_total_mercancias_exentas" xml:"ResumenFactura>TotalMercanciasExentas"`
	ResumenTotalMercExonerada      decimal.Decimal         `ts:"decimal,null" sql:"type:decimal(18,5);" json:"resumen_total_merc_exonerada" xml:"ResumenFactura>TotalMercExonerada"`
	ResumenTotalGravado            decimal.Decimal         `ts:"decimal,null" sql:"type:decimal(18,5);" json:"resumen_total_gravado" xml:"ResumenFactura>TotalGravado"`
	ResumenTotalExento             decimal.Decimal         `ts:"decimal,null" sql:"type:decimal(18,5);" json:"resumen_total_exento" xml:"ResumenFactura>TotalExento"`
	ResumenTotalExonerado          decimal.Decimal         `ts:"decimal,null" sql:"type:decimal(18,5);" json:"resumen_total_exonerado" xml:"ResumenFactura>TotalExonerado"`
	ResumenTotalVenta              decimal.Decimal         `ts:"decimal,null" gorm:"not null" sql:"type:decimal(18,5);" json:"resumen_total_venta" xml:"ResumenFactura>TotalVenta"`
	ResumenTotalDescuentos         decimal.Decimal         `ts:"decimal,null" sql:"type:decimal(18,5);" json:"resumen_total_descuentos" xml:"ResumenFactura>TotalDescuentos"`
	ResumenTotalVentaNeta          decimal.Decimal         `ts:"decimal,null" gorm:"not null" sql:"type:decimal(18,5);" json:"resumen_total_venta_neta" xml:"ResumenFactura>TotalVentaNeta"`
	ResumenTotalImpuesto           decimal.Decimal         `ts:"decimal,null" sql:"type:decimal(18,5);" json:"resumen_total_impuesto" xml:"ResumenFactura>TotalImpuesto"`
	ResumenTotalIVADevuelto        decimal.Decimal         `ts:"decimal,null" sql:"type:decimal(18,5);" json:"resumen_total_iva_devuelto" xml:"ResumenFactura>TotalIVADevuelto"`
	ResumenTotalOtrosCargos        decimal.Decimal         `ts:"decimal,null" sql:"type:decimal(18,5);" json:"resumen_total_otros_cargos" xml:"ResumenFactura>TotalOtrosCargos"`
	ResumenTotalComprobante        decimal.Decimal         `ts:"decimal,null" gorm:"not null" sql:"type:decimal(18,5);" json:"resumen_total_comprobante" xml:"ResumenFactura>TotalComprobante"`
	InformacionReferencia          []InformacionReferencia `ts:",null" json:"informacion_referencia" xml:"InformacionReferencia"`
	NormativaNumeroResolucion      string                  `gorm:"not null" json:"normativa_numero_resolucion" xml:"Normativa>NumeroResolucion"`
	NormativaFechaResolucion       model.CustomTime        `ts:"date,null" gorm:"not null" json:"normativa_fecha_resolucion" xml:"Normativa>FechaResolucion"`
}

type Exoneracion struct {
	gorm.Model
	ImpuestoID            uint            `ts:"-"`
	TipoDocumento         int             `json:"tipo_documento" xml:"Exoneracion>TipoDocumento"`
	NumeroDocumento       string          `json:"numero_documento" xml:"Exoneracion>NumeroDocumento"`
	NombreInstitucion     string          `json:"nombre_institucion" xml:"Exoneracion>NombreInstitucion"`
	FechaEmision          time.Time       `json:"fecha_emision" xml:"Exoneracion>FechaEmision"`
	PorcentajeExoneracion decimal.Decimal `ts:"decimal,null" sql:"type:decimal(4,2)"`
	MontoExoneracion      decimal.Decimal `ts:"decimal,null" sql:"type:decimal(18,5)"`
}

// Impuesto representa los impuestos en la factura
type Impuesto struct {
	gorm.Model
	Codigo         int             `gorm:"not null" validate:"required" json:"codigo"`
	CodigoTarifa   int             `json:"codigo_tarifa"`
	Tarifa         decimal.Decimal `ts:"decimal,null" gorm:"not null" validate:"required" sql:"type:decimal(18,5);" json:"tarifa"`
	FactorIVA      decimal.Decimal `ts:"decimal,null" json:"factor_iva" sql:"type:decimal(18,5)"`
	Monto          decimal.Decimal `ts:"decimal,null" gorm:"not null" validate:"required" sql:"type:decimal(18,5);" json:"monto"`
	Exoneracion    Exoneracion     `ts:",null" json:"exoneracion"`
	LineaDetalleID uint            `ts:"-" json:"linea_detalle_id"`
}

// InformacionReferencia representa una referencia a otros comprobantes electronicos.
type InformacionReferencia struct {
	gorm.Model
	TipoDoc      int              `json:"tipo_doc" xml:"TipoDoc"`
	Numero       string           `json:"numero" xml:"Numero"`
	FechaEmision model.CustomTime `ts:"date,null" json:"fecha_emision" xml:"FechaEmision"`
	Codigo       int              `json:"codigo" xml:"Codigo"`
	Razon        string           `json:"razon" xml:"Razon"`
	DocumentoID  uint             `ts:"-" json:"documento_id"`
}

type CodigoComercial struct {
	gorm.Model
	Tipo           int    `ts:"CodigoComercialTipo" validate:"oneof=1 2 3 4 99" json:"tipo"`
	Codigo         string `validate:"max=20" json:"codigo"`
	LineaDetalleID uint   `ts:"-" json:"linea_detalle_id"`
}

// LineaDetalle representa los productos y servicios que se agregan a la factura
type LineaDetalle struct {
	gorm.Model
	NumeroLinea           int               `gorm:"not null" validate:"required,numeric,min=1" json:"numero_linea"`
	PartidaArancelaria    string            `json:"partida_arancelaria"`
	CodigosComerciales    []CodigoComercial `json:"codigos_comerciales"`
	CodigoProducto        string            `json:"codigo_producto" validate:"max=13"`
	Cantidad              decimal.Decimal   `ts:"decimal,null" gorm:"not null" sql:"type:decimal(16,3);" json:"cantidad" validate:"required,numeric,max=9999999999999999.999"`
	UnidadMedida          string            `gorm:"not null" validate:"required" json:"unidad_medida"`
	UnidadMedidaComercial string            `json:"unidad_medida_comercial"`
	Detalle               string            `gorm:"not null" validate:"required,max=160" json:"detalle"`
	PrecioUnitario        decimal.Decimal   `ts:"decimal" gorm:"not null" validate:"required,numeric,max=9999999999999.99999" sql:"type:decimal(18,5);" json:"precio_unitario"`
	MontoTotal            decimal.Decimal   `ts:"decimal" gorm:"not null" validate:"required,numeric,max=9999999999999.99999" sql:"type:decimal(18,5);" json:"monto_total"`
	MontoDescuento        decimal.Decimal   `ts:"decimal,null" sql:"type:decimal(18,5);" json:"monto_descuento" validate:"omitempty,numeric,max=9999999999999.99999" xml:"Descuento>MontoDescuento"`
	NaturalezaDescuento   string            `json:"naturaleza_descuento" xml:"Descuento>NaturalezaDescuento"`
	SubTotal              decimal.Decimal   `ts:"decimal" gorm:"not null" validate:"required,numeric,max=9999999999999.99999" sql:"type:decimal(18,5);" json:"sub_total"`
	BaseImponible         decimal.Decimal   `ts:"decimal,null" sql:"type:decimal(18,5)"`
	Impuesto              []Impuesto        `ts:",null" json:"impuesto" validate:"omitempty" xml:"Impuesto"`
	ImpuestoNeto          decimal.Decimal   `ts:"decimal,null" json:"impuesto_neto" sql:"type:decimal(18,5);"`
	MontoTotalLinea       decimal.Decimal   `ts:"decimal" gorm:"not null" validate:"required,numeric,max=9999999999999.99999" sql:"type:decimal(18,5);" json:"monto_total_linea"`
	DocumentoID           uint              `ts:"-" json:"documento_id"`
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
	GrupoID                    uint            `ts:"-" json:"grupo_id"`
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
	MontoTotalImpuesto         decimal.Decimal `ts:"decimal,null" sql:"type:decimal(18,5);" json:"monto_total_impuesto"`
	TotalFactura               decimal.Decimal `ts:"decimal,null" sql:"type:decimal(18,5);" json:"total_factura"`
}

// Receptor representa a un receptor de comprobantes electrónicos, es similar
// al tipo Emisor con la diferencia de que este tipo añade un campo adicional
type Receptor struct {
	gorm.Model
	Nombre                   string                   `json:"nombre"`
	IdentificacionTipo       model.IdentificacionTipo `json:"identificacion_tipo" xml:"Identificacion>Tipo"`
	IdentificacionNumero     string                   `gorm:"unique" json:"identificacion_numero" xml:"Identificacion>Numero"`
	NombreComercial          string                   `json:"nombre_comercial"`
	Provincia                int                      `gorm:"not null" json:"provincia" validate:"numeric,oneof=1 2 3 4 5 6 7" xml:"Ubicacion>Provincia"`
	Canton                   int                      `gorm:"not null" json:"canton" validate:"numeric,required" xml:"Ubicacion>Canton"`
	Distrito                 int                      `gorm:"not null" json:"distrito" validate:"numeric,required" xml:"Ubicacion>Distrito"`
	Barrio                   int                      `json:"barrio" validate:"numeric,required" xml:"Ubicacion>Barrio"`
	OtrasSenas               string                   `gorm:"size:250" json:"otras_senas" validate:"required" xml:"Ubicacion>OtrasSenas"`
	TelCodigoPais            int                      `json:"tel_codigo_pais" validate:"numeric" xml:"Telefono>CodigoPais"`
	TelNumTelefono           string                   `gorm:"size:20" json:"tel_num_telefono" xml:"Telefono>NumTelefono"`
	FaxCodigoPais            int                      `json:"fax_codigo_pais" validate:"numeric,omitempty" xml:"Fax>CodigoPais"`
	FaxNumTelefono           string                   `gorm:"size:20" json:"fax_num_telefono" xml:"Fax>NumTelefono"`
	CorreoElectronico        string                   `json:"correo_electronico"`
	IdentificacionExtranjero string                   `json:"identificacion_extranjero"`
	OtrasSenasExtranjero     string                   `json:"otras_senas_extranjero" validate:"max=300"`
}

// Usuario representa los datos de autenticación de los usuarios de la
// aplicación web
type Usuario struct {
	gorm.Model
	SuscripcionID int       `ts:"-" json:"suscripcion_id"`
	Cuenta        string    `gorm:"not null;unique" validate:"email" json:"cuenta"`
	Clave         string    `gorm:"not null" json:"-"`
	Plano         string    `gorm:"-" validate:"min=8" json:"plano"`
	Permisos      []Permiso `gorm:"many2many:usuario_permisos;"`
}

// BeforeSave es llamado por GORM antes de escribir el record en la base de
// datos. Comprueba el tamaño de la contraseña y la cifra usando BCrypt.
func (u *Usuario) BeforeSave(db *gorm.DB) error {
	validate := validator.New()
	// Valida los datos de entrada
	err := validate.Struct(u)
	if err != nil {
		return err
	}
	// Hace un hash de la contraseña
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Plano), 14)
	if err != nil {
		return err
	}
	u.Clave = string(bytes[:])
	return nil
}

// OAuthResponse respresenta la respuesta que da el servidor OAuth de Hacienda
type OAuthResponse struct {
	ID               uint      `ts:"-" gorm:"primary_key" json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	AccessToken      string    `json:"access_token"`
	ExpiresIn        int       `json:"expires_in"`
	RefreshExpiresIn int       `json:"refresh_expires_in"`
	RefreshToken     string    `json:"refresh_token"`
}

// Credencial guarda información de autenticación de un obligado
// contribuyente procedientes de Hacienda
type Credencial struct {
	gorm.Model
	InvalidAt  time.Time `json:"invalido_desde"`
	Usuario    string    `gorm:"not null;unique" json:"usuario"`
	Clave      string    `gorm:"not null" json:"clave"`
	Pin        string    `gorm:"not null" json:"pin"`
	ArchivoP12 string    `gorm:"not null;unique" json:"archivo_p_12"`
	Contenido  string    `gorm:"-" json:"contenido"`
}

// Grupo conserva información para la emisión de facturas
type Grupo struct {
	gorm.Model
	IdentificacionTipo   model.IdentificacionTipo `gorm:"not null" json:"identificacion_tipo" validate:"numeric"`
	NombreComercial      string                   `gorm:"size:80" json:"nombre_comercial"`
	Provincia            int                      `gorm:"not null" json:"provincia" validate:"numeric,oneof=1 2 3 4 5 6 7"`
	Canton               int                      `gorm:"not null" json:"canton" validate:"numeric,required"`
	Distrito             int                      `gorm:"not null" json:"distrito" validate:"numeric,required"`
	Barrio               int                      `json:"barrio" validate:"numeric,required"`
	OtrasSenas           string                   `gorm:"size:250" json:"otras_senas" validate:"required"`
	TelCodigoPais        int                      `json:"tel_codigo_pais" validate:"numeric"`
	TelNumTelefono       string                   `gorm:"size:20" json:"tel_num_telefono"`
	FaxCodigoPais        int                      `json:"fax_codigo_pais" validate:"numeric,omitempty"`
	FaxNumTelefono       string                   `gorm:"size:20" json:"fax_num_telefono"`
	CorreoElectronico    string                   `gorm:"not null;size:60" json:"correo_electronico" validate:"email"`
	Nombre               string                   `gorm:"not null;size:100" json:"nombre"`
	IdentificacionNumero string                   `gorm:"not null;size:14" json:"identificacion_numero"`
	Actividades          []*Actividad             `json:"actividades" gorm:"many2many:ga"`
	OAuth                OAuthResponse            `gorm:"foreignkey:OAuthResponseID" json:"o_auth"`
	Credenciales         Credencial               `gorm:"foreignkey:CredencialID" json:"credenciales"`
	CodigoContent        string                   `gorm:"-" json:"codigo_content"`
	Codigos              []Codigo                 `json:"codigos"`
	OAuthResponseID      uint                     `ts:"-" json:"o_auth_response_id"`
	CredencialID         uint                     `ts:"-" json:"credencial_id"`
	SuscripcionID        uint                     `ts:"-" json:"suscripcion_id"`
	Documentos           []Documento              `json:"documentos"`
	Mensajes             []MensajeHacienda        `json:"mensajes"`
}

type Emisor struct {
	gorm.Model
	IdentificacionTipo   model.IdentificacionTipo `gorm:"not null" json:"identificacion_tipo" validate:"numeric"`
	NombreComercial      string                   `gorm:"size:80" json:"nombre_comercial"`
	Provincia            int                      `gorm:"not null" json:"provincia" validate:"numeric,oneof=1 2 3 4 5 6 7" xml:"Ubicacion>Provincia"`
	Canton               int                      `gorm:"not null" json:"canton" validate:"numeric,required" xml:"Ubicacion>Canton"`
	Distrito             int                      `gorm:"not null" json:"distrito" validate:"numeric,required" xml:"Ubicacion>Distrito"`
	Barrio               int                      `json:"barrio" validate:"numeric,required" xml:"Ubicacion>Barrio"`
	OtrasSenas           string                   `gorm:"size:250" json:"otras_senas" validate:"required" xml:"Ubicacion>OtrasSenas"`
	TelCodigoPais        int                      `json:"tel_codigo_pais" validate:"numeric" xml:"Telefono>CodigoPais"`
	TelNumTelefono       string                   `gorm:"size:20" json:"tel_num_telefono" xml:"Telefono>NumTelefono"`
	FaxCodigoPais        int                      `json:"fax_codigo_pais" validate:"numeric,omitempty" xml:"Fax>CodigoPais"`
	FaxNumTelefono       string                   `gorm:"size:20" json:"fax_num_telefono" xml:"Fax>NumTelefono"`
	CorreoElectronico    string                   `gorm:"not null;size:60" json:"correo_electronico" validate:"email"`
	Nombre               string                   `gorm:"not null;size:100" json:"nombre"`
	IdentificacionNumero string                   `gorm:"not null;size:14" json:"identificacion_numero"`
}

// Suscripcion representa una suscripción al servicio de generación de
// comprobantes electrónicos, también contiene permisos de acceso para
// distintas areas de la aplicación web
type Suscripcion struct {
	gorm.Model
	PaidAt   time.Time             `json:"paid_at"`
	Tipo     model.SuscripcionTipo `gorm:"type:int" json:"tipo"`
	Usuarios []Usuario             `json:"usuarios"`
	Grupo    Grupo                 `json:"grupo"`
}

// Codigo es una tabla que enlaza distintos códigos con detalles
type Codigo struct {
	gorm.Model
	Codigo  string `gorm:"not null;size:80;uniqueIndex:idx_codigo_grupoid" json:"codigo"`
	Detalle string `gorm:"not null;size:160" json:"detalle"`
	GrupoID uint   `ts:"-" json:"grupo_id" gorm:"uniqueIndex:idx_codigo_grupoid"`
}

// Recepcion sirve para manejar todos los documentos que clientes reciben en
// virtud de ser receptores de comprobantes electrónicos
type Recepcion struct {
	Documento
	EstadoRecepcion model.EstadoReceptorTipo `json:"estado_recepcion"`
}

// Consecutivo mantiene una lista de consecutivos para emisores y receptores de
// comprobantes electrónicos
type Consecutivo struct {
	gorm.Model
	GrupoID  uint                  `ts:"-" json:"grupo_id" gorm:"uniqueIndex:idx_grupo_tipo_caja_sucursal"`
	Tipo     model.ConsecutivoTipo `json:"tipo" validate:"required,oneof=1 2 3 4 5" gorm:"uniqueIndex:idx_grupo_tipo_caja_sucursal"`
	Sucursal int                   `json:"sucursal" validate:"required,min=1,max=999" gorm:"uniqueIndex:idx_grupo_tipo_caja_sucursal"`
	Caja     int                   `json:"caja" validate:"required,min=1,max=99999" gorm:"uniqueIndex:idx_grupo_tipo_caja_sucursal"`
	Contador int64                 `json:"contador" gorm:"default:'1'" validate:"required,min=0,max=9999999999"`
}

type OtrosCargos struct {
	gorm.Model
	TipoDocumento int             `validate:"oneof=1 2 3 4 5 6 7 99,required" json:"tipo_documento"`
	Detalle       string          `validate:"max=160,required" json:"detalle"`
	Porcentaje    decimal.Decimal `ts:"decimal,null" sql:"type:decimal(4,2)" json:"porcentaje"`
	MontoCargo    decimal.Decimal `ts:"decimal,null" sql:"type:decimal(18,5)" json:"monto_cargo"`
	DocumentoID   uint            `ts:"-" json:"documento_id"`
}

type Actividad struct {
	gorm.Model
	Codigo      string   `gorm:"unique;not null;size=6" json:"codigo"`
	Descripcion string   `gorm:"not null" json:"descripcion"`
	Grupos      []*Grupo `ts:"-" gorm:"many2many:ga" json:"-"`
}

// Canton es una tabla de cantones
type Canton struct {
	ID        int    `ts:"-" json:"id" gorm:"primary_key"`
	Provincia int    `json:"provincia" gorm:"not null"`
	Canton    int    `json:"canton" gorm:"not null"`
	Nombre    string `json:"nombre" gorm:"not null"`
}

// Distrito es una tabla de distritos, su relación con la tabla de Canton es
// implícita por el ID de Canton
type Distrito struct {
	ID        int    `ts:"-" json:"id" gorm:"primary_key"`
	Provincia int    `json:"provincia" gorm:"not null"`
	Canton    int    `json:"canton" gorm:"not null"`
	Distrito  int    `json:"distrito" gorm:"not null"`
	Nombre    string `json:"nombre" gorm:"not null"`
}

// Barrio es una tabla de barrios, su relación con la tabla de Barrio es
// implícita por el ID de distrito
type Barrio struct {
	ID        int    `ts:"-" json:"id" gorm:"primary_key"`
	Provincia int    `json:"provincia" gorm:"not null"`
	Canton    int    `json:"canton" gorm:"not null"`
	Distrito  int    `json:"distrito" gorm:"not null"`
	Barrio    int    `json:"barrio" gorm:"not null"`
	Nombre    string `json:"nombre" gorm:"not null"`
}

// Permiso define acciones que pueden realizarse sobre un sujeto
type Permiso struct {
	ID        uint       `ts:"-" gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `ts:"-" json:"-"`
	UpdatedAt time.Time  `ts:"-" json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Sujeto    string     `gorm:"unique_index:idx_sujeto_accion" json:"sujeto"`
	Accion    string     `gorm:"unique_index:idx_sujeto_accion" json:"accion"`
}
