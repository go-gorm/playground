#!/bin/bash

go mod tidy

echo Testing with gorm v1.23.9
go get gorm.io/gorm@v1.23.9
go get gorm.io/driver/sqlite@v1.3.6
go mod tidy
go run . gorm-v1.23.9

GORM=v1.24.6
echo Testing with gorm $GORM
go get gorm.io/driver/sqlite@v1.4.0
go get gorm.io/gorm@$GORM
go mod tidy
go run . gorm-$GORM

for i in `seq 0 6`; do
  GORM=v1.25.$i
  echo Testing with gorm $GORM
  go get gorm.io/driver/sqlite@v1.4.0
  go get gorm.io/gorm@$GORM
  go mod tidy
  go run . gorm-$GORM
done

rm *.db
