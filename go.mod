module gorm.io/playground

go 1.16

require (
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/jackc/pgconn v1.14.1 // indirect
	github.com/jackc/pgx/v4 v4.18.1 // indirect
	github.com/jackc/pgx/v5 v5.4.2 // indirect
	github.com/microsoft/go-mssqldb v1.4.0 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	gorm.io/driver/mysql v1.5.1
	gorm.io/driver/postgres v1.5.2
	gorm.io/driver/sqlite v1.5.2
	gorm.io/driver/sqlserver v1.5.1
	gorm.io/gorm v1.25.2-0.20230530020048-26663ab9bf55
)

replace gorm.io/gorm => ./gorm
