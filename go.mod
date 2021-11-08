module gorm.io/playground

go 1.16

require (
	github.com/go-kit/kit v0.10.0 // indirect
	golang.org/x/crypto v0.0.0-20211108221036-ceb1ce70b4fa // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gorm.io/driver/mysql v1.1.3
	gorm.io/driver/postgres v1.2.1
	gorm.io/driver/sqlite v1.2.3
	gorm.io/driver/sqlserver v1.2.0
	gorm.io/gorm v1.22.2
)

replace gorm.io/gorm => ./gorm
