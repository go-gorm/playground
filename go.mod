module gorm.io/playground

go 1.16

require (
	github.com/denisenkom/go-mssqldb v0.12.2 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	gorm.io/driver/mysql v1.4.1
	gorm.io/driver/postgres v1.4.4
	gorm.io/driver/sqlite v1.4.2
	gorm.io/driver/sqlserver v1.4.1
	gorm.io/gorm v1.24.5
)

replace gorm.io/gorm => ./gorm
