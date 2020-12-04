module gorm.io/playground

go 1.14

require (
	github.com/stretchr/testify v1.6.1
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
	gorm.io/driver/mysql v1.0.3
	gorm.io/driver/postgres v1.0.5
	gorm.io/driver/sqlite v1.1.4
	gorm.io/driver/sqlserver v1.0.5
	gorm.io/gorm v1.20.8
)

replace gorm.io/gorm => ./gorm
