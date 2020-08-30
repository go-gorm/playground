module gorm.io/playground

go 1.14

require (
	github.com/go-playground/validator/v10 v10.3.0
	github.com/lib/pq v1.7.1
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shopspring/decimal v1.2.0
	github.com/sirupsen/logrus v1.6.0
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	gorm.io/driver/postgres v1.0.0
	gorm.io/gorm v1.9.19
)

replace gorm.io/gorm => ./gorm
