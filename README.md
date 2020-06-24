# GORM PlayGround

GORM PlayGround can be used to play GORM and reports issues

[![test status](https://github.com/go-gorm/playground/workflows/tests/badge.svg?branch=master "test status")](https://github.com/go-gorm/playground/actions)

### Usage

Modify [https://github.com/go-gorm/playground/edit/master/main_test.go](https://github.com/go-gorm/playground/edit/master/main_test.go) and create pull report, your code will runs with sqlite, mysql, postgres and sqlserver with the help of Github Action

If you encounter a bug in GORM, please report it at [https://github.com/go-gorm/gorm/issues](https://github.com/go-gorm/gorm/issues) with the PlayGround Pull Request's link

Your code will run with GORM's lastest master branch, if you want to use a different one, please change following comment in the above [link](https://github.com/go-gorm/playground/edit/master/main_test.go) to your favorite one

```go
// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
```

BTW, we have prepared some structs with relationships in [https://github.com/go-gorm/playground/blob/master/models.go](https://github.com/go-gorm/playground/blob/master/models.go) that you may interested to use

Don't forgot to close your PR after finish play! ;)

## Happy Hacking!
