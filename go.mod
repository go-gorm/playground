module gorm.io/playground

go 1.16

require (
	github.com/denisenkom/go-mssqldb v0.12.2 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/jackc/pgx/v4 v4.18.1 // indirect
	github.com/jackc/pgx/v5 v5.3.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	github.com/microsoft/go-mssqldb v0.21.0 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	gorm.io/driver/mysql v1.4.7
	gorm.io/driver/postgres v1.5.0
	gorm.io/driver/sqlite v1.4.4
	gorm.io/driver/sqlserver v1.4.2
	gorm.io/gorm v1.24.7-0.20230306060331-85eaf9eeda11
)

replace gorm.io/gorm => ./gorm
