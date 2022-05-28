module gorm.io/playground

go 1.16

require (
	github.com/denisenkom/go-mssqldb v0.12.2 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/mattn/go-sqlite3 v1.14.13 // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	gorm.io/driver/mysql v1.3.3
	gorm.io/driver/sqlite v1.3.2
	gorm.io/driver/sqlserver v1.3.2
	gorm.io/gorm v1.23.5
)

// Issue does not happen with this version set
require (
	//gorm.io/driver/postgres v1.3.5
)

// Issue starts to happen with this version set
require (
	gorm.io/driver/postgres v1.3.6
)

//replace gorm.io/gorm => ./gorm
