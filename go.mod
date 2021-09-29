module gorm.io/playground

go 1.16

require (
	github.com/denisenkom/go-mssqldb v0.10.0 // indirect
	github.com/google/go-cmp v0.3.1
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/mattn/go-sqlite3 v1.14.7 // indirect
	golang.org/x/crypto v0.0.0-20210505212654-3497b51f5e64 // indirect
	golang.org/x/text v0.3.6 // indirect
	gorm.io/driver/mysql v1.0.6
	gorm.io/driver/postgres v1.1.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/driver/sqlserver v1.0.7
	gorm.io/gorm v1.21.9
	gotest.tools v2.2.0+incompatible
)

replace gorm.io/gorm => ./gorm
