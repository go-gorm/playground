module gorm.io/playground

go 1.16

require (
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/hashicorp/go-secure-stdlib/base62 v0.1.2 // indirect
	github.com/jackc/pgx/v4 v4.14.0 // indirect
	golang.org/x/crypto v0.0.0-20211117183948-ae814b36b871 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gorm.io/driver/mysql v1.2.0
	gorm.io/driver/postgres v1.2.2
	gorm.io/driver/sqlite v1.2.6
	gorm.io/driver/sqlserver v1.2.1
	gorm.io/gorm v1.22.3
)

replace gorm.io/gorm => ./gorm
