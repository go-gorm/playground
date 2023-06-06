module gorm.io/playground

go 1.16

require (
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/microsoft/go-mssqldb v1.1.0 // indirect
	gorm.io/driver/mysql v1.5.1
	gorm.io/driver/postgres v1.5.2
	gorm.io/driver/sqlite v1.5.1
	gorm.io/driver/sqlserver v1.5.0
	gorm.io/gorm v1.25.1
)

replace gorm.io/gorm => ./gorm
