module github.com/thanhpp/playground

go 1.16

require (
	github.com/jackc/pgconn v1.10.1
	github.com/jackc/pgerrcode v0.0.0-20201024163028-a0d42d470451
	github.com/jackc/pgproto3/v2 v2.2.0 // indirect
	github.com/jackc/pgtype v1.9.0 // indirect
	github.com/jinzhu/now v1.1.3 // indirect
	golang.org/x/crypto v0.0.0-20211117183948-ae814b36b871 // indirect
	gorm.io/driver/mysql v1.2.0
	gorm.io/driver/postgres v1.2.2
	gorm.io/driver/sqlite v1.2.4
	gorm.io/driver/sqlserver v1.2.1
	gorm.io/gorm v1.22.3
)

replace gorm.io/gorm => ./gorm
