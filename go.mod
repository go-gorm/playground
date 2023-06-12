module gorm.io/playground

go 1.16

require (
	gorm.io/driver/mysql v1.4.1
	gorm.io/driver/postgres v1.5.2
	gorm.io/driver/sqlite v1.4.2
	gorm.io/driver/sqlserver v1.4.1
	gorm.io/gorm v1.25.1
)

//replace gorm.io/gorm => ./gorm
