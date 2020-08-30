#!/bin/bash -e

dialects=("sqlite" "mysql" "postgres" "sqlserver" "oracle")
ORACLE_INSTANT_CLIENT_URL="https://download.oracle.com/otn_software/linux/instantclient/19600/instantclient-basic-linux.x64-19.6.0.0.0dbru.zip"
ORACLE_INSTANT_CLIENT_FILE="instant_client.zip"

if [ "$GORM_ENABLE_CACHE" = "" ]
then
rm -rf gorm
fi

[ -d gorm ] || (echo "git clone --depth 1 -b $(cat main_test.go | grep GORM_BRANCH | awk '{print $3}') $(cat main_test.go | grep GORM_REPO | awk '{print $3}')"; git clone --depth 1 -b $(cat main_test.go | grep GORM_BRANCH | awk '{print $3}') $(cat main_test.go | grep GORM_REPO | awk '{print $3}'))

cp go.mod go.mod.bak
sed '/gorm.io\/driver/d' go.mod.bak > go.mod

for dialect in "${dialects[@]}" ; do
  if [ "$GORM_DIALECT" = "" ] || [ "$GORM_DIALECT" = "${dialect}" ]
  then
    if [[ $(grep TEST_DRIVER main_test.go) =~ "${dialect}" ]]
    then
      if [[ "$dialect" =~ "oracle" ]]
      then
        if [[ ! -d $(pwd)/instantclient_19_6 ]]
        then
          if [[ ! -f "$ORACLE_INSTANT_CLIENT_FILE" ]]
          then
            echo "downloading oracle instant client..."
            curl "$ORACLE_INSTANT_CLIENT_URL" -o "$ORACLE_INSTANT_CLIENT_FILE"
          fi
          echo "unzipping oracle instant client..."
          unzip -o "$ORACLE_INSTANT_CLIENT_FILE"
        fi
        export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:$(pwd)/instantclient_19_6
        echo "exported instant client libraries to LD_LIBRARY_PATH, now it should not complain about missing oracle libraries"
      fi
      echo "testing ${dialect}..."
      GORM_DIALECT=${dialect} go test -race -count=1 -v ./...
    else
      echo "skip ${dialect}..."
    fi
  fi
done

mv go.mod.bak go.mod
