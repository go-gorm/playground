module gorm.io/playground

go 1.16

require (
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.9 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gorm.io/driver/mysql v1.1.2
	gorm.io/driver/postgres v1.1.2
	gorm.io/driver/sqlite v1.1.6
	gorm.io/driver/sqlserver v1.1.0
	gorm.io/gorm v1.21.16
)

replace gorm.io/gorm => ./gorm
