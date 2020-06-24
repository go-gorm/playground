# GORM PlayGround

GORM PlayGround can be used to play GORM and reports issues

### Usage

Modify [https://github.com/go-gorm/playground/blob/master/main_test.go](https://github.com/go-gorm/playground/blob/master/main_test.go) and create pull report, your code will be run automatically with Github Action

If you encounter a bug in GORM, please report it at [https://github.com/go-gorm/gorm/issues](https://github.com/go-gorm/gorm/issues) with the PlayGround Pull Request's link

Your code will run with GORM's lastest master branch, if you want to use a different one, please change following comment in the above [link](https://github.com/go-gorm/playground/blob/master/main_test.go) to your favorite one

```go
// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
```

Don't forgot to close your PR after finish play! ;)

## Happy Hacking!
