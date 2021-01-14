# GORM Playground

GORM Playground can be used to play GORM and reports issues, if you encounter a bug in GORM, please report it at [https://github.com/go-gorm/gorm/issues](https://github.com/go-gorm/gorm/issues) with the Playground Pull Request's link

[![test status](https://github.com/go-gorm/playground/workflows/tests/badge.svg?branch=master "test status")](https://github.com/go-gorm/playground/actions)

### Quick Start

##### 1. [Fork this repo](https://docs.github.com/en/free-pro-team@latest/github/getting-started-with-github/fork-a-repo)

##### 2. [Clone the forked repo to your local](https://docs.github.com/en/free-pro-team@latest/github/creating-cloning-and-archiving-repositories/cloning-a-repository)

##### 3. Setup test database

```bash
# install docker-compose https://docs.docker.com/compose/install/

# setup test databases
docker-compose up
```

##### 4. Run tests with lastest GORM and all drivers

```bash
./test.sh

# Run tests with cached GORM and latest drivers
GORM_ENABLE_CACHE=true ./test.sh

# Run tests with specfied database
GORM_DIALECT=mysql go test
```

##### 5. Modify tests and make it fail

##### 6. [Create Playground Pull Request](https://docs.github.com/en/free-pro-team@latest/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request) and [Create a GORM issue](https://github.com/go-gorm/gorm/issues/new?template=bug_report.md) with the link

### Advanced Usage

We are using the following configuration run your code (GORM's latest master branch, latest database drivers: sqlite, mysql, postgres, sqlserver), you could change the configuration in file [main_test.go](https://github.com/go-gorm/playground/edit/master/main_test.go)

```go
// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver
```

We have prepared some structs with relationships in [https://github.com/go-gorm/playground/blob/master/models.go](https://github.com/go-gorm/playground/blob/master/models.go) that you can use for your tests

## Happy Hacking!
