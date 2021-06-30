module gorm.io/playground

go 1.16

require (
	github.com/denisenkom/go-mssqldb v0.10.0 // indirect
	github.com/jackc/pgproto3/v2 v2.1.0 // indirect
	github.com/jackc/pgx/v4 v4.11.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.7 // indirect
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	gorm.io/driver/mysql v1.1.1
	gorm.io/driver/postgres v1.1.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/driver/sqlserver v1.0.7
	gorm.io/gorm v1.21.11
)

replace gorm.io/gorm => ./gorm
