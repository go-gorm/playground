#!/bin/bash -e

dialects=("sqlserver")

if [ "$GORM_ENABLE_CACHE" = "" ]
then
rm -rf gorm
fi

[ -d gorm ] || (echo "git clone --depth 1 -b $(cat main_test.go | grep GORM_BRANCH | awk '{print $3}') $(cat main_test.go | grep GORM_REPO | awk '{print $3}')"; git clone --depth 1 -b $(cat main_test.go | grep GORM_BRANCH | awk '{print $3}') $(cat main_test.go | grep GORM_REPO | awk '{print $3}'))

go get -u -t ./...

# SqlServer for Mac M1
if [[ -z $GITHUB_ACTION ]]; then
  if [[ $(uname -a) == *" arm64" ]]; then
    MSSQL_IMAGE=mcr.microsoft.com/azure-sql-edge docker compose up --detach --quiet-pull || true
    echo "starting"
    go install github.com/microsoft/go-sqlcmd/cmd/sqlcmd@latest || true
    SQLCMDPASSWORD=LoremIpsum86 sqlcmd -U sa -S localhost:9930 -Q "IF DB_ID('gorm') IS NULL CREATE DATABASE gorm" > /dev/null || true
    SQLCMDPASSWORD=LoremIpsum86 sqlcmd -U sa -S localhost:9930 -Q "IF SUSER_ID (N'gorm') IS NULL CREATE LOGIN gorm WITH PASSWORD = 'LoremIpsum86';" > /dev/null || true
    SQLCMDPASSWORD=LoremIpsum86 sqlcmd -U sa -S localhost:9930 -Q "IF USER_ID (N'gorm') IS NULL CREATE USER gorm FROM LOGIN gorm; ALTER SERVER ROLE sysadmin ADD MEMBER [gorm];" > /dev/null || true
  else
    docker compose up --detach --quiet-pull
    echo "starting..."
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

docker compose down
