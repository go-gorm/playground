module gorm.io/playground

go 1.16

require (
	github.com/stretchr/testify v1.8.4
	gorm.io/driver/sqlite v1.4.2
	gorm.io/gorm v1.24.0
)

//replace gorm.io/gorm => ./gorm
