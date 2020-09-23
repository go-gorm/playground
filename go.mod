module gorm.io/playground

go 1.14

require (
	gorm.io/datatypes v0.0.0-20200806042100-bc394008dd0d
	gorm.io/driver/mysql v0.3.1
	gorm.io/driver/postgres v0.2.6
	gorm.io/driver/sqlite v1.0.8
	gorm.io/driver/sqlserver v0.2.5
	gorm.io/gorm v1.20.1
)

replace gorm.io/gorm => ./gorm
