package model

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

// XSDATETIME es un formato de fecha y hora que Hacienda espera
var XSDATETIME = "2006-01-02T15:04:05"

// NMDATETIME es un formato de fecha y hora que Hacienda espera que se presente
// el valor time.Time de la resolución
var NMDATETIME = "2006-01-02 15:04:05"

// DocumentoTipo define un nuevo tipo para diferenciar documentos
type DocumentoTipo int

const (
	// FACTURA se refiere a Factura Electronica
	FACTURA DocumentoTipo = iota + 1
	// DEBITO se refiere a Nota de Débito Electronica
	DEBITO
	// CREDITO se refiere a Nota de Crédito Electronica
	CREDITO
	// TIQUETE se refiere a Tiquete Electrónico
	TIQUETE
	FACTURAEXPO
	FACTURACOM
)

func (d DocumentoTipo) String() string {
	name := "Comprobante electronico"
	switch d {
	case FACTURA:
		name = "Factura Electronica"
	case DEBITO:
		name = "Nota de Débito Electronica"
	case CREDITO:
		name = "Nota de Crédito Electronica"
	case TIQUETE:
		name = "Tiquete Electronico"
	case FACTURAEXPO:
		name = "Factura Electronica Exportacion"
	case FACTURACOM:
		name = "Factura Electronica Compra"
	}
	return name
}

// SuscripcionTipo Define el tipo de suscripción del cliente, el calculo de el
// día de cobro se calcula a partir del día de creación de la suscripción.
type SuscripcionTipo int

const (
	// DESACTIVADA indica que la suscripción no esta activa
	DESACTIVADA SuscripcionTipo = iota
	// MENSUAL indica que la suscripción se cobra por mes
	MENSUAL
	// ANUAL indica que la suscripción se cobra por año
	ANUAL
	// CORTECIA indica que la suscripción es de cortecía :)
	CORTECIA
	// Trimestral indica que la suscripción es trimestral
	GRATUITA
)

func (s SuscripcionTipo) String() string {
	name := "Suscripcion de tipo desconocido"
	switch s {
	case DESACTIVADA:
		name = "Desactivada"
	case MENSUAL:
		name = "Mensual"
	case ANUAL:
		name = "Anual"
	case CORTECIA:
		name = "Cortecia"
	case GRATUITA:
		name = "Gratuita"
	}
	return name
}

// EstadoTipo se usa para indicar el tipo de estado en el cual se encuentra el
// comprobante electrónico
type EstadoTipo int

const (
	// REPOSO dice que la factura solo fue creada pero aun no se manda a
	// Hacienda
	REPOSO EstadoTipo = iota
	// ENVIADO es para cuando el comprobante fue enviado
	ENVIADO
	// RECIBIDO es para cuando el comprobante fue recibido por Hacienda
	RECIBIDO
	// PROCESADO dice que Hacienda ha procesado el comprobante
	PROCESADO
	// ERROR dice que algo malo ocurrió de lado de Hacienda (?)
	ERROR
	// ACEPTADO es cuando Hacienda no esta fallando y de hecho el
	// comprobante fue aceptado
	ACEPTADO
	// RECHAZADO dice que el comprobante fue rechazado por Hacienda
	RECHAZADO
	// ErrorInterno habla sobre un error interno en Fero, este valor no
	// debe ser grabado con el comprobante en la base de datos.
	ErrorInterno
)

func (t EstadoTipo) String() (tipo string) {
	switch t {
	case REPOSO:
		tipo = "Reposo"
	case ENVIADO:
		tipo = "Enviado"
	case RECIBIDO:
		tipo = "Recibido"
	case PROCESADO:
		tipo = "Procesado"
	case ERROR:
		tipo = "Error"
	case ACEPTADO:
		tipo = "Aceptado"
	case RECHAZADO:
		tipo = "Rechazado"
	}
	return tipo
}

type IdentificacionTipo int

const (
	PersonaFisica IdentificacionTipo = iota + 1
	PersonaJuridica
	NITE
	DIMEX
)

func (i IdentificacionTipo) String() (tipo string) {
	switch i {
	case PersonaFisica:
		tipo = "Cédula persona física"
	case PersonaJuridica:
		tipo = "Cédula persona jurídica"
	case NITE:
		tipo = "NITE"
	case DIMEX:
		tipo = "DIMEX"
	}
	return
}

// EstadoReceptorTipo
type EstadoReceptorTipo int

const (
	RPREVISAR EstadoReceptorTipo = iota
	RPACEPTADO
	RPACEPTADOPARCIAL
	RPRECHAZADO
)

type ConsecutivoTipo int

const (
	CReceptor ConsecutivoTipo = iota + 1
	CFactura
	CCredito
	CDebito
	CTiquete
	CFACTURACOM
	CFACTURAEXPO
)

var (
	UNO  = decimal.NewFromFloat(1.0)
	CIEN = decimal.NewFromFloat(100.0)
)

type CustomTime struct {
	time.Time
}

func (c *CustomTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	var parse time.Time
	switch start.Name.Local {
	case "FechaEmision":
		parse, err = time.Parse(XSDATETIME, v)
		if err != nil {
			// Algunos emisores especifican el Huso horario
			parse, err = time.Parse(XSDATETIME+"-07:00", v)
		}
	case "FechaResolucion":
		parse, err = time.Parse(NMDATETIME, v)
	}
	if err != nil {
		return fmt.Errorf("%s: %s", start.Name.Local, err)
	}
	*c = CustomTime{parse}
	return nil
}

func (c *CustomTime) Scan(v interface{}) error {
	vb, ok := v.(time.Time)
	if ok {
		*c = CustomTime{vb}
	}
	return nil
}

func (c CustomTime) Value() (driver.Value, error) {
	return c.Time, nil
}
