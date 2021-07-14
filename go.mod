module gorm.io/playground

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/denisenkom/go-mssqldb v0.10.0 // indirect
	github.com/jackc/pgx/v4 v4.12.0 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	gorm.io/driver/mysql v1.1.1
	gorm.io/driver/postgres v1.1.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/driver/sqlserver v1.0.7
	gorm.io/gorm v1.21.12
)

replace gorm.io/gorm => ./gorm
