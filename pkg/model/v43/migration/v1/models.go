package v1

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Documento struct {
	gorm.Model
	CodigoActividad           string `gorm:"not null;size:6"`
	OtrosCargos               []OtrosCargos
	ResumenTotalServExonerado decimal.Decimal `sql:"type:decimal(18,5);"`
	ResumenTotalMercExonerada decimal.Decimal `sql:"type:decimal(18,5);"`
	ResumenTotalExonerado     decimal.Decimal `sql:"type:decimal(18,5);"`
	ResumenTotalIVADevuelto   decimal.Decimal `sql:"type:decimal(18,5)"`
	ResumenTotalOtrosCargos   decimal.Decimal `sql:"type:decimal(18,5);"`
}

type LineaDetalle struct {
	gorm.Model
	NumeroLinea           int `gorm:"not null"`
	PartidaArancelaria    string
	CodigosComerciales    []CodigoComercial
	CodigoProducto        string
	Cantidad              decimal.Decimal `gorm:"not null" sql:"type:decimal(16,3);"`
	UnidadMedida          string          `gorm:"not null"`
	UnidadMedidaComercial string
	Detalle               string          `gorm:"not null" `
	PrecioUnitario        decimal.Decimal `gorm:"not null" sql:"type:decimal(18,5);"`
	MontoTotal            decimal.Decimal `gorm:"not null" sql:"type:decimal(18,5);"`
	MontoDescuento        decimal.Decimal `sql:"type:decimal(18,5);"`
	NaturalezaDescuento   string
	SubTotal              decimal.Decimal `gorm:"not null" sql:"type:decimal(18,5);"`
	BaseImponible         decimal.Decimal `sql:"type:decimal(18,5)"`
	Impuesto              []Impuesto
	ImpuestoNeto          decimal.Decimal ` sql:"type:decimal(18,5);"`
	MontoTotalLinea       decimal.Decimal `gorm:"not null" sql:"type:decimal(18,5);"`
	DocumentoID           uint
}

type Impuesto struct {
	gorm.Model
	Codigo         int `gorm:"not null"`
	CodigoTarifa   int
	Tarifa         decimal.Decimal `gorm:"not null" sql:"type:decimal(18,5);" `
	FactorIVA      decimal.Decimal ` sql:"type:decimal(18,5)"`
	Monto          decimal.Decimal `gorm:"not null" sql:"type:decimal(18,5);"`
	Exoneracion    Exoneracion
	LineaDetalleID uint
}

type Exoneracion struct {
	gorm.Model
	ImpuestoID            uint
	TipoDocumento         int
	NumeroDocumento       string
	NombreInstitucion     string
	FechaEmision          time.Time
	PorcentajeExoneracion decimal.Decimal `sql:"type:decimal(3,2)"`
	MontoExoneracion      decimal.Decimal `sql:"type:decimal(18,5)"`
}

type Receptor struct {
	OtrasSenasExtranjero string ` gorm:"type:text" `
}

type Grupo struct {
	Nombre     string `gorm:"size:100" `
	OtrasSenas string `gorm:"size:250" `
}

type OtrosCargos struct {
	gorm.Model
	TipoDocumento int
	Detalle       string
	Porcentaje    decimal.Decimal `sql:"type:decimal(3,2)"`
	MontoCargo    decimal.Decimal `sql:"type:decimal(18,5)"`
	DocumentoID   uint
}

type CodigoComercial struct {
	gorm.Model
	Tipo           int
	Codigo         string
	LineaDetalleID uint
}
