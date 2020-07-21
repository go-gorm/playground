module gorm.io/playground

go 1.14

require (
	github.com/godror/godror v0.17.1 // indirect
	github.com/stevefan1999-personal/gorm-driver-oracle v0.0.0-20200720134927-d43380d395ca
	gorm.io/driver/mysql v0.3.0
	gorm.io/driver/postgres v0.2.6
	gorm.io/driver/sqlite v1.0.8
	gorm.io/driver/sqlserver v0.2.4
	gorm.io/gorm v0.2.23
)

replace gorm.io/gorm => ./gorm
