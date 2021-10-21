module gorm.io/playground

go 1.16

require (
	github.com/jackc/pgx/v4 v4.13.0
	github.com/mattn/go-sqlite3 v1.14.9 // indirect
	gorm.io/driver/mysql v1.1.2
	gorm.io/driver/postgres v1.1.2
	gorm.io/driver/sqlite v1.1.6
	gorm.io/driver/sqlserver v1.1.0
	gorm.io/gorm v1.21.16
)

replace gorm.io/gorm => ./gorm
