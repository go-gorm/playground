package v9

type Actividad struct{}
type Barrio struct{}
type Canton struct{}
type Codigo struct{}
type CodigoComercial struct{}
type Consecutivo struct{}
type Credencial struct{}
type Distrito struct{}
type Documento struct{}
type Emisor struct{}
type Exoneracion struct{}
type Grupo struct{}
type Ga struct{}              // usado por Grupo en many2many
type UsuarioPermisos struct{} // usado en many2many
type Impuesto struct{}
type InformacionReferencia struct{}
type LineaDetalle struct{}
type MensajeHacienda struct{}
type OAuthResponse struct{}
type OtrosCargos struct{}
type Permiso struct{}
type Recepcion struct{}
type Receptor struct{}
type Suscripcion struct{}
type Usuario struct{}

const (
	actividad              = "actividades"
	barrio                 = "barrios"
	canton                 = "cantones"
	codigo                 = "codigos"
	codigo_comercial       = "codigo_comercials"
	consecutivo            = "consecutivos"
	credencial             = "credencials"
	distrito               = "distritos"
	documento              = "documentos"
	emisor                 = "emisores"
	exoneracion            = "exoneracions"
	grupo                  = "grupos"
	ga                     = "ga"
	usuario_permisos       = "usuario_permisos"
	impuesto               = "impuestos"
	informacion_referencia = "informacion_referencia"
	linea_detalle          = "linea_detalles"
	mensaje_hacienda       = "mensaje_haciendas"
	o_auth_response        = "o_auth_responses"
	otros_cargos           = "otros_cargos"
	permiso                = "permisos"
	recepcion              = "recepcions"
	receptor               = "receptors"
	suscripcion            = "suscripciones"
	usuario                = "usuarios"
)
