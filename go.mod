module gorm.io/playground

go 1.22.0

toolchain go1.23.2

require (
	gorm.io/driver/mysql v1.5.7
	gorm.io/driver/postgres v1.5.9
	gorm.io/driver/sqlite v1.5.6
	gorm.io/driver/sqlserver v1.5.3
	gorm.io/gen v0.3.26
	gorm.io/gen/examples v0.0.0-00010101000000-000000000000
	gorm.io/gorm v1.25.12
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.1 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.24 // indirect
	github.com/microsoft/go-mssqldb v1.7.2 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	golang.org/x/tools v0.26.0 // indirect
	gorm.io/datatypes v1.2.4 // indirect
	gorm.io/hints v1.1.2 // indirect
	gorm.io/plugin/dbresolver v1.5.3 // indirect
)

replace gorm.io/gorm => ./gorm

replace gorm.io/gen/examples => ./gen/examples

// these allow testing against the last version where the tests passed (v1.25.5)
replace gorm.io/driver/sqlite => gorm.io/driver/sqlite v1.5.4

replace gorm.io/driver/postgres => gorm.io/driver/postgres v1.5.4

replace gorm.io/driver/mysql => gorm.io/driver/mysql v1.5.2

replace gorm.io/driver/sqlserver => gorm.io/driver/sqlserver v1.5.2
