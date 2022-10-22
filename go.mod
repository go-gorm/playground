module gorm.io/playground

go 1.16

require (
	golang.org/x/crypto v0.1.0 // indirect
	gorm.io/driver/mysql v1.4.3
	gorm.io/driver/postgres v1.4.5
	gorm.io/driver/sqlite v1.4.3
	gorm.io/driver/sqlserver v1.4.1
	gorm.io/gorm v1.24.1-0.20221019064659-5dd2bb482755
)

replace gorm.io/gorm => ./gorm
