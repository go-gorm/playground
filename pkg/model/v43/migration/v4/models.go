package v4

import (
	"github.com/shopspring/decimal"
)

type Exoneracion struct {
	PorcentajeExoneracion decimal.Decimal `sql:"type:decimal(4,2)"`
}

type OtrosCargos struct {
	Porcentaje decimal.Decimal `sql:"type:decimal(4,2)"`
}
