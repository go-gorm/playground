module gorm.io/playground

go 1.16

require (
	github.com/denisenkom/go-mssqldb v0.10.0 // indirect
	github.com/jackc/pgx/v4 v4.13.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.8 // indirect
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210813211128-0a44fdfbc16e // indirect
	golang.org/x/text v0.3.7 // indirect
	gorm.io/driver/mysql v1.1.2
	gorm.io/driver/postgres v1.1.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/driver/sqlserver v1.0.7
	gorm.io/gorm v1.21.13
)

replace gorm.io/gorm => ./gorm
