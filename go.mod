module gorm.io/playground

go 1.16

require (
	github.com/go-kit/kit v0.10.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gorm.io/driver/mysql v1.1.3
	gorm.io/driver/postgres v1.2.1
	gorm.io/driver/sqlite v1.2.2
	gorm.io/driver/sqlserver v1.1.2
	gorm.io/gorm v1.22.1
	gorm.io/plugin/dbresolver v1.1.0 // indirect
)

replace gorm.io/gorm => ./gorm
