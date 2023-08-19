module gorm.io/playground

go 1.16

require (
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/microsoft/go-mssqldb v1.5.0 // indirect
	golang.org/x/tools v0.12.0 // indirect
	gorm.io/datatypes v1.2.0 // indirect
	gorm.io/driver/mysql v1.5.1
	gorm.io/driver/postgres v1.5.2
	gorm.io/driver/sqlite v1.5.3 // indirect
	gorm.io/driver/sqlserver v1.5.1
	gorm.io/gen v0.3.23
	gorm.io/gorm v1.25.2
	gorm.io/hints v1.1.2 // indirect
	gorm.io/plugin/dbresolver v1.4.6 // indirect
)

replace gorm.io/gorm => ./gorm
