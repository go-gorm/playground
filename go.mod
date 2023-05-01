module gorm.io/playground

go 1.16

require (
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/jackc/pgx/v5 v5.3.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	golang.org/x/crypto v0.8.0 // indirect
	gorm.io/driver/mysql v1.5.0
	gorm.io/driver/postgres v1.5.0
	gorm.io/driver/sqlite v1.5.0
	gorm.io/driver/sqlserver v1.4.3
	gorm.io/gorm v1.25.0
)

replace gorm.io/gorm => ./gorm
