module gorm.io/playground

go 1.14

require (
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/google/uuid v1.1.1
	gorm.io/driver/mysql v0.3.0
	gorm.io/driver/postgres v0.2.6
	gorm.io/driver/sqlite v1.0.8
	gorm.io/driver/sqlserver v0.2.4
	gorm.io/gorm v0.2.23
)

replace gorm.io/gorm => ./gorm
