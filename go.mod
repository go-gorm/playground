module gorm.io/playground

go 1.16

require (
	github.com/denisenkom/go-mssqldb v0.10.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.7 // indirect
	gorm.io/driver/mysql v1.0.6
	gorm.io/driver/postgres v1.2.2
	gorm.io/driver/sqlite v1.1.4
	gorm.io/driver/sqlserver v1.0.7
	gorm.io/gorm v1.22.2
)

replace gorm.io/gorm => ./gorm
