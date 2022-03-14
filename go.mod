module gorm.io/playground

go 1.16

require (
	github.com/jackc/pgx/v4 v4.15.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.11 // indirect
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292 // indirect
	gorm.io/driver/mysql v1.3.2
	gorm.io/driver/postgres v1.3.1
	gorm.io/driver/sqlite v1.3.1
	gorm.io/driver/sqlserver v1.3.1
	gorm.io/gorm v1.23.2
)

// replace gorm.io/gorm => ./gorm
