module gorm.io/playground

go 1.14

require (
	github.com/mattn/go-sqlite3 v1.14.5 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/driver/postgres v1.0.5
	gorm.io/driver/sqlite v1.1.3
	gorm.io/driver/sqlserver v1.0.5
	gorm.io/gorm v1.20.7
)

replace gorm.io/gorm => ./gorm
