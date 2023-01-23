module gorm.io/playground

go 1.16

require (
	github.com/denisenkom/go-mssqldb v0.12.2 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/jackc/pgtype v1.13.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	golang.org/x/crypto v0.5.0 // indirect
	gorm.io/driver/mysql v1.4.5
	gorm.io/driver/postgres v1.4.6
	gorm.io/driver/sqlite v1.4.4
	gorm.io/driver/sqlserver v1.4.2
	gorm.io/gorm v1.24.3
)

replace gorm.io/gorm => ./gorm
