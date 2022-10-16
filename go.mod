module gorm.io/playground

go 1.16

require (
	github.com/denisenkom/go-mssqldb v0.12.2 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	golang.org/x/crypto v0.0.0-20221012134737-56aed061732a // indirect
	golang.org/x/text v0.3.8 // indirect
	gorm.io/driver/mysql v1.4.1
	gorm.io/driver/postgres v1.4.4
	gorm.io/driver/sqlite v1.4.2
	gorm.io/driver/sqlserver v1.4.1
	gorm.io/gorm v1.24.0
)

replace gorm.io/gorm => ./gorm
