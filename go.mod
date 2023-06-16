module gorm.io/playground

go 1.16

require (
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/jackc/pgx/v5 v5.4.0 // indirect
	golang.org/x/crypto v0.10.0 // indirect
	gorm.io/driver/mysql v1.5.1
	gorm.io/driver/postgres v1.5.2
	gorm.io/driver/sqlite v1.5.2
	gorm.io/driver/sqlserver v1.5.1
	gorm.io/gorm v1.25.2-0.20230530020048-26663ab9bf55
)

replace gorm.io/gorm => ./gorm
