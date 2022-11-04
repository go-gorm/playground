module gorm.io/playground

go 1.16

require (
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	golang.org/x/crypto v0.1.0 // indirect
	gorm.io/driver/mysql v1.4.3
	gorm.io/driver/postgres v1.4.5
	gorm.io/driver/sqlite v1.4.3
	gorm.io/driver/sqlserver v1.4.1
	gorm.io/gorm v1.24.1
)

replace gorm.io/gorm => ./gorm
