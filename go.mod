module gorm.io/playground

go 1.16

replace gorm.io/driver/sqlite => github.com/Arsen6331/sqlite v1.3.3-0.20220429065518-98334794f73c

require (
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/jackc/pgx/v4 v4.16.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/crypto v0.0.0-20220511200225-c6db032c6c88 // indirect
	golang.org/x/sys v0.0.0-20220503163025-988cb79eb6c6 // indirect
	golang.org/x/tools v0.1.10 // indirect
	golang.org/x/xerrors v0.0.0-20220411194840-2f41105eb62f // indirect
	gorm.io/driver/mysql v1.3.3
	gorm.io/driver/postgres v1.3.5
	gorm.io/driver/sqlite v1.3.2
	gorm.io/driver/sqlserver v1.3.2
	gorm.io/gorm v1.23.5
	lukechampine.com/uint128 v1.2.0 // indirect
	modernc.org/libc v1.16.8 // indirect
	modernc.org/opt v0.1.3 // indirect
	modernc.org/sqlite v1.17.2 // indirect
)

replace gorm.io/gorm => ./gorm
