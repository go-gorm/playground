module gorm.io/playground

go 1.14

require (
	gorm.io/driver/mysql v1.0.4
	gorm.io/driver/postgres v1.0.7
	gorm.io/driver/sqlite v1.1.4
	gorm.io/driver/sqlserver v1.0.6
	gorm.io/gorm v1.20.12
)

replace gorm.io/gorm => ./gorm
