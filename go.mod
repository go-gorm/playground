module gorm.io/playground

go 1.16

require (
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/stretchr/testify v1.8.0
	gorm.io/driver/mysql v1.3.3
	gorm.io/driver/postgres v1.3.10
	gorm.io/driver/sqlite v1.3.2
	gorm.io/driver/sqlserver v1.3.2
	gorm.io/gorm v1.23.10
)

//replace gorm.io/gorm => ./gorm
