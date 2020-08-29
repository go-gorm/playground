module gorm.io/playground

go 1.14

require (
	github.com/stevefan1999-personal/gorm-driver-oracle v0.0.0-20200829094520-df84a6d6b3cd
	gorm.io/driver/mysql v1.0.0
	gorm.io/driver/postgres v1.0.0
	gorm.io/driver/sqlite v1.1.0
	gorm.io/driver/sqlserver v1.0.0
	gorm.io/gorm v1.9.19
)

replace gorm.io/gorm => ./gorm