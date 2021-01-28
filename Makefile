
test-mysql:
	GORM_ENABLE_CACHE=true GORM_DIALECT=mysql go test

test-sqlite:
	GORM_ENABLE_CACHE=true GORM_DIALECT=sqlite go test

test-postgres:
	GORM_ENABLE_CACHE=true GORM_DIALECT=postgres go test
