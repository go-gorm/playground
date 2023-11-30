#!/bin/bash -e

dialects=("sqlite" "mysql" "postgres")

if [ "$GORM_ENABLE_CACHE" = "" ]
then
rm -rf gorm
fi

[ -d gorm ] || (echo "git clone --depth 1 -b $(cat main_test.go | grep GORM_BRANCH | awk '{print $3}') $(cat main_test.go | grep GORM_REPO | awk '{print $3}')"; git clone --depth 1 -b $(cat main_test.go | grep GORM_BRANCH | awk '{print $3}') $(cat main_test.go | grep GORM_REPO | awk '{print $3}'))

go get -u -t ./...

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
