module gorm.io/playground

go 1.16

require (
	github.com/jackc/pgx/v4 v4.14.1 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	golang.org/x/crypto v0.0.0-20211202192323-5770296d904e // indirect
	gorm.io/driver/mysql v1.2.1
	gorm.io/driver/postgres v1.2.3
	gorm.io/driver/sqlite v1.2.6
	gorm.io/driver/sqlserver v1.2.1
	gorm.io/gorm v1.22.4
)

replace gorm.io/gorm => ./gorm
