module gorm.io/playground

go 1.14

require (

	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/driver/postgres v1.0.5
	gorm.io/driver/sqlite v1.1.3
	gorm.io/driver/sqlserver v1.0.4
	gorm.io/gorm v1.20.5
)

//replace gorm.io/gorm => ./gorm
