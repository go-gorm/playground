module gorm.io/playground

go 1.14

require (
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kr/pty v1.1.8 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	golang.org/x/tools v0.0.0-20191029190741-b9c20aec41a5 // indirect
	gorm.io/driver/mysql v1.0.5
	gorm.io/driver/postgres v1.0.8
	gorm.io/driver/sqlite v1.1.4
	gorm.io/driver/sqlserver v1.0.7
	gorm.io/gorm v1.21.4
)

replace gorm.io/gorm => ./gorm
