module gorm.io/playground

go 1.14

require (
	github.com/stevefan1999-personal/gorm-driver-oracle v0.0.0-20200719072112-4b23102fa303
	gorm.io/driver/mysql v0.2.9
	gorm.io/driver/postgres v0.2.5
	gorm.io/driver/sqlite v1.0.8
	gorm.io/driver/sqlserver v0.2.4
	gorm.io/gorm v0.2.22
)

replace gorm.io/gorm => ./gorm
