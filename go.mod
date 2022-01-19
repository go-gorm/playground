module gorm.io/playground

go 1.16

require (
	github.com/jackc/pgx/v4 v4.14.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.10 // indirect
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce // indirect
	gorm.io/driver/mysql v1.2.3
	gorm.io/driver/postgres v1.2.3
	gorm.io/driver/sqlite v1.2.6
	gorm.io/driver/sqlserver v1.2.1
	gorm.io/gorm v1.22.4
)

replace gorm.io/gorm => ./gorm
