package models43

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lib/pq"
	"github.com/shopspring/decimal"
	"gitlab.com/shackra/gormbug/pkg/model"
	"gorm.io/gorm"
)

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func randomString(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		panic(fmt.Sprintf("error al generar bytes al azar: %s", err))
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes)
}

func randomNumber(n int) string {
	var letter = []rune("0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func PopulateDB(db *gorm.DB) error {
	rand.Seed(19902812)
	var superError error
	// popula la base de datos
	for subi := 0; subi < 3; subi++ {
		var permisos []Permiso
		var nombre string
		if subi == 0 {
			permisos = []Permiso{
				{Sujeto: "invoice", Accion: "create"},
				{Sujeto: "invoice", Accion: "view"},
				{Sujeto: "admin", Accion: "create"},
				{Sujeto: "admin", Accion: "read"},
				{Sujeto: "admin", Accion: "update"},
				{Sujeto: "admin", Accion: "change"},
				{Sujeto: "admin", Accion: "view"},
				{Sujeto: "issuer", Accion: "change"},
				{Sujeto: "mh_credentials", Accion: "change"},
				{Sujeto: "mh_credentials", Accion: "read"},
				{Sujeto: "user_account", Accion: "change"},
				{Sujeto: "user_account", Accion: "read"},
			}
			nombre = randomString(30) + " admin"
		} else {
			nombre = randomString(30)
			permisos = []Permiso{
				{Sujeto: "invoice", Accion: "create"},
				{Sujeto: "invoice", Accion: "view"},
				{Sujeto: "issuer", Accion: "change"},
				{Sujeto: "mh_credentials", Accion: "change"},
				{Sujeto: "mh_credentials", Accion: "read"},
				{Sujeto: "user_account", Accion: "change"},
				{Sujeto: "user_account", Accion: "read"},
			}
		}
		sub := Suscripcion{
			Tipo: 1,
			Usuarios: []Usuario{
				{
					Cuenta:   randomString(5) + "@" + randomString(5) + ".dev",
					Plano:    "holamundo",
					Permisos: permisos,
				},
			},
		}
		var actividades []*Actividad
		err := db.Find(&actividades, "id = 1 OR id = 2 OR id = 3").Error
		if err != nil {
			return err
		}
		grp := Grupo{
			Nombre:               nombre,
			IdentificacionTipo:   1,
			IdentificacionNumero: randomNumber(9),
			NombreComercial:      "Comer. " + nombre,
			Provincia:            1,
			Canton:               1,
			Distrito:             1,
			Barrio:               1,
			Actividades:          actividades,
			CorreoElectronico:    randomString(5) + "@" + randomString(5) + ".dev",
			OAuth: OAuthResponse{
				AccessToken:      randomString(10),
				RefreshToken:     randomString(20),
				ExpiresIn:        10,
				RefreshExpiresIn: 10,
			},
			Credenciales: Credencial{
				Usuario:    randomString(6) + "@" + randomString(5) + ".dev",
				Clave:      randomString(5),
				Pin:        randomNumber(4),
				ArchivoP12: randomString(1),
			},
		}
		// coloca documentos dentro de este grupo
		for doci := 0; doci < 10; doci++ {
			rec := Receptor{
				Nombre:               randomString(25),
				IdentificacionTipo:   1,
				IdentificacionNumero: randomNumber(9),
				NombreComercial:      "Comer. " + randomString(10),
				Provincia:            1,
				Canton:               1,
				Distrito:             1,
				Barrio:               1,
				OtrasSenas:           randomString(10),
				TelCodigoPais:        506,
				TelNumTelefono:       randomNumber(8),
				CorreoElectronico:    randomString(5) + "@" + randomString(5) + ".dev",
			}
			year, month, day := time.Now().Date()
			// Este arreglo para tener el año en dos digitos debería durar 100 años
			year = year - 2000
			consecutivo := fmt.Sprintf("%03d%05d%02d%010d", 1, 1, model.FACTURA, doci)
			doc := Documento{
				Tipo:                           model.FACTURA,
				Estado:                         model.RECHAZADO,
				Clave:                          fmt.Sprintf("506%02d%02d%02d%012s%s1%08d", day, month, year, grp.IdentificacionNumero, consecutivo, rand.Intn(99999999)),
				NumeroConsecutivo:              consecutivo,
				FechaEmision:                   model.CustomTime{Time: time.Now()},
				Receptor:                       rec,
				CondicionVenta:                 1,
				PlazoCredito:                   1,
				MedioPago:                      pq.Int64Array{1},
				ResumenCodigoMoneda:            "CRC",
				ResumenTipoCambio:              decimal.NewFromFloat(1.0),
				ResumenTotalServGravados:       decimal.NewFromFloat(10.0),
				ResumenTotalServExentos:        decimal.NewFromFloat(10.0),
				ResumenTotalMercanciasGravadas: decimal.NewFromFloat(10.0),
				ResumenTotalMercanciasExentas:  decimal.NewFromFloat(10.0),
				ResumenTotalGravado:            decimal.NewFromFloat(10.0),
				ResumenTotalExento:             decimal.NewFromFloat(10.0),
				ResumenTotalVenta:              decimal.NewFromFloat(10.0),
				ResumenTotalDescuentos:         decimal.NewFromFloat(10.0),
				ResumenTotalVentaNeta:          decimal.NewFromFloat(10.0),
				ResumenTotalImpuesto:           decimal.NewFromFloat(10.0),
				ResumenTotalComprobante:        decimal.NewFromFloat(10.0),
			}
			for lineai := 0; lineai < 10; lineai++ {
				linea := LineaDetalle{
					NumeroLinea:        lineai + 1,
					CodigosComerciales: []CodigoComercial{{Tipo: 1, Codigo: "0000"}},
					Cantidad:           decimal.NewFromFloat(1),
					UnidadMedida:       "Sp",
					Detalle:            "cosa para pruebas",
					PrecioUnitario:     decimal.NewFromFloat(1),
					MontoTotal:         decimal.NewFromFloat(1),
					SubTotal:           decimal.NewFromFloat(1),
					Impuesto: []Impuesto{
						{
							Codigo: 7,
							Tarifa: decimal.NewFromFloat(10),
							Monto:  decimal.NewFromFloat(1),
						}, {
							Codigo: 1,
							Tarifa: decimal.NewFromFloat(13),
							Monto:  decimal.NewFromFloat(1),
							Exoneracion: Exoneracion{
								TipoDocumento:         1,
								NumeroDocumento:       "001100",
								NombreInstitucion:     "Institución inexistente",
								FechaEmision:          time.Now(),
								MontoExoneracion:      decimal.NewFromFloat(1),
								PorcentajeExoneracion: decimal.NewFromFloat(5),
							},
						},
					},
					MontoTotalLinea: decimal.NewFromFloat(1),
				}
				// inserta la linea de detalle al documento
				doc.LineaDetalle = append(doc.LineaDetalle, linea)
			}
			// inserta el documento creado a la lista de documentos
			// del grupo
			grp.Documentos = append(grp.Documentos, doc)
		}
		// insertamos el grupo a la suscripción
		sub.Grupo = grp
		superError = db.Create(&sub).Error
		if superError != nil {
			break
		}
	}
	return superError
}

func FindConsecutivo(db *gorm.DB, gid uint, sucursal, caja int, t model.ConsecutivoTipo) Consecutivo {
	var consecutivo Consecutivo
	db.Where(Consecutivo{GrupoID: gid, Sucursal: sucursal, Caja: caja, Tipo: t}).FirstOrCreate(&consecutivo)
	if consecutivo.Contador < 1 {
		consecutivo.Contador = 1
	}
	return consecutivo
}
