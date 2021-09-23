module gorm.io/playground

go 1.16

require (
	github.com/denisenkom/go-mssqldb v0.11.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gorm.io/datatypes v1.0.2
	gorm.io/driver/mysql v1.1.2
	gorm.io/driver/postgres v1.1.1
	gorm.io/driver/sqlite v1.1.5
	gorm.io/driver/sqlserver v1.0.9
	gorm.io/gorm v1.21.15
)

replace gorm.io/gorm => ./gorm
