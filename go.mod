module gorm.io/playground

go 1.16

require (
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/jackc/pgx/v5 v5.4.2 // indirect
	github.com/microsoft/go-mssqldb v1.4.0 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	gorm.io/driver/mysql v1.5.1
	gorm.io/driver/postgres v1.5.2
	gorm.io/driver/sqlite v1.5.2
	gorm.io/driver/sqlserver v1.5.1
	gorm.io/gorm v1.25.2
)

// Issue does not happen with this version set
require (
	//gorm.io/plugin/dbresolver v1.4.1
)

// Issue starts to happen with this version set
require (
	gorm.io/plugin/dbresolver v1.4.2
)

//replace gorm.io/gorm => ./gorm
