module gorm.io/playground

go 1.16

require (
	github.com/denisenkom/go-mssqldb v0.12.2 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/jackc/pgx/v4 v4.16.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.13 // indirect
	golang.org/x/crypto v0.0.0-20220518034528-6f7dac969898 // indirect
	gorm.io/driver/mysql v1.3.3
	gorm.io/driver/postgres v1.3.5
	gorm.io/driver/sqlite v1.3.2
	gorm.io/driver/sqlserver v1.3.2
	gorm.io/gorm v1.23.4
)

replace gorm.io/gorm => ./gorm
