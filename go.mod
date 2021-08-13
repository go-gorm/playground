module gorm.io/playground

go 1.16

require (
	github.com/denisenkom/go-mssqldb v0.10.0 // indirect
	github.com/google/uuid v1.0.0
	github.com/jackc/pgx/v4 v4.13.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.8 // indirect
	golang.org/x/crypto v0.0.0-20210812204632-0ba0e8f03122 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gorm.io/driver/mysql v1.1.2
	gorm.io/driver/postgres v1.1.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/driver/sqlserver v1.0.7
	gorm.io/gorm v1.21.12
)

replace gorm.io/gorm => ./gorm
