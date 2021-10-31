module gorm.io/playground

go 1.16

require (
	gorm.io/driver/mysql v1.1.3
	gorm.io/driver/postgres v1.2.1
	gorm.io/driver/sqlite v1.2.3
	gorm.io/driver/sqlserver v1.1.2
	gorm.io/gorm v1.22.2
)

replace gorm.io/gorm => ./gorm
