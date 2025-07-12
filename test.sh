#!/bin/bash -e

dialects=("sqlite" "mysql" "postgres")

if [ "$GORM_ENABLE_CACHE" = "" ]
then
rm -rf gorm
fi

[ -d gorm ] || (echo "git clone --depth 1 -b $(grep GORM_BRANCH main_test.go | awk '{print $3}') $(grep GORM_REPO main_test.go | awk '{print $3}')"; git clone --depth 1 -b $(grep GORM_BRANCH main_test.go | awk '{print $3}') $(grep GORM_REPO main_test.go | awk '{print $3}'))

go get -u -t ./...


# SqlServer for Mac M1
if [[ -z $GITHUB_ACTION ]]; then
  if [ -d tests ]
  then
    cd tests
    if [[ $(uname -a) == *" arm64" ]]; then
      MSSQL_IMAGE=mcr.microsoft.com/azure-sql-edge docker compose up -d --wait
      go install github.com/microsoft/go-sqlcmd/cmd/sqlcmd@latest
      for query in \
        "IF DB_ID('gorm') IS NULL CREATE DATABASE gorm" \
        "IF SUSER_ID (N'gorm') IS NULL CREATE LOGIN gorm WITH PASSWORD = '${SQLCMDPASSWORD}';" \
        "IF USER_ID (N'gorm') IS NULL CREATE USER gorm FROM LOGIN gorm; ALTER SERVER ROLE sysadmin ADD MEMBER [gorm];"
      do
        SQLCMDPASSWORD=LoremIpsum86 sqlcmd -U sa -S localhost:9930 -Q "$query" > /dev/null
      done
    else
      MSSQL_IMAGE=mcr.microsoft.com/mssql/server docker compose up -d --wait
    fi
    cd ..
  fi
fi

for dialect in "${dialects[@]}" ; do
  if [ "$GORM_DIALECT" = "" ] || [ "$GORM_DIALECT" = "${dialect}" ]
  then
    if [[ $(grep TEST_DRIVER main_test.go) =~ "${dialect}" ]]
    then
      echo "testing ${dialect}..."
      GORM_DIALECT=${dialect} go test -race -count=1 -v ./...
    else
      echo "skip ${dialect}..."
    fi
  fi
done
