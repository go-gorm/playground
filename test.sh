#!/bin/sh -e

dialects=("sqlite" "mysql" "postgres" "sqlserver")

rm -rf gorm

echo "git clone --depth 1 -b $(cat main_test.go | grep GORM_BRANCH | awk '{print $3}') $(cat main_test.go | grep GORM_REPO | awk '{print $3}')"

git clone --depth 1 -b $(cat main_test.go | grep GORM_BRANCH | awk '{print $3}') $(cat main_test.go | grep GORM_REPO | awk '{print $3}')

if [ -d tests ]
then
  cd tests
  cp go.mod go.mod.bak
  sed '/$[[:space:]]*gorm.io\/driver/d' go.mod.bak > go.mod
  cd ..
fi

for dialect in "${dialects[@]}" ; do
  if [ "$GORM_DIALECT" = "" ] || [ "$GORM_DIALECT" = "${dialect}" ]
  then
    echo "testing ${dialect}..."
    GORM_DIALECT=${dialect} go test -race -count=1 -v ./...
  fi
done

if [ -d tests ]
then
  cd tests
  mv go.mod.bak go.mod
fi
