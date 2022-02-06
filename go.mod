module gorm.io/playground

go 1.16

require (
	github.com/denisenkom/go-mssqldb v0.12.0 // indirect
	github.com/jackc/pgx/v4 v4.14.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.11 // indirect
	golang.org/x/crypto v0.0.0-20220131195533-30dcbda58838 // indirect
	gorm.io/driver/mysql v1.2.3
	gorm.io/driver/postgres v1.2.3
	gorm.io/driver/sqlite v1.2.6
	gorm.io/driver/sqlserver v1.2.1
	gorm.io/gorm v1.22.5
)

replace gorm.io/gorm => ./gorm
