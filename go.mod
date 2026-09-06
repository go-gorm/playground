module gorm.io/playground

go 1.26.0

require (
	github.com/stretchr/testify v1.12.1
	gorm.io/driver/mysql v1.6.0
	gorm.io/driver/postgres v1.6.2
	gorm.io/driver/sqlite v1.6.0
	gorm.io/driver/sqlserver v1.6.4
	gorm.io/gorm v1.31.2
)

require (
	filippo.io/edwards25519 v1.2.0 // indirect
	github.com/go-sql-driver/mysql v1.10.1 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.10.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.52 // indirect
	github.com/microsoft/go-mssqldb v1.11.0 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/crypto v0.56.0 // indirect
	golang.org/x/sync v0.22.0 // indirect
	golang.org/x/text v0.41.0 // indirect
)

replace gorm.io/gorm => ./gorm

// replace gorm.io/driver/postgres => ../postgres
