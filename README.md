# GORM PlayGround

GORM PlayGround can be used to play GORM and reports issues

[![test status](https://github.com/go-gorm/playground/workflows/tests/badge.svg?branch=master "test status")](https://github.com/go-gorm/playground/actions)

### Usage

Modify [https://github.com/go-gorm/playground/edit/master/main_test.go](https://github.com/go-gorm/playground/edit/master/main_test.go) and create pull report, your code will runs with sqlite, mysql, postgres and sqlserver with the help of Github Action

If you encounter a bug in GORM, please report it at [https://github.com/go-gorm/gorm/issues](https://github.com/go-gorm/gorm/issues) with the PlayGround Pull Request's link

We are using following configuration run your code (GORM's latest master branch, database drivers: sqlite, mysql, postgres, sqlserver), you could change it in the above [link](https://github.com/go-gorm/playground/edit/master/main_test.go)

```go
// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver
```

BTW, we have prepared some structs with relationships in [https://github.com/go-gorm/playground/blob/master/models.go](https://github.com/go-gorm/playground/blob/master/models.go) that you may interested to use

Don't forgot to close your PR after finish play! ;)

## Happy Hacking!
