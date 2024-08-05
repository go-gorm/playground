package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

/*
NewDB 创建数据库连接
*/
func main() { //nolint:gocritic // only used in init
	var (
		err            error
		resolverConfig dbresolver.Config
	)
	dsn := "root:123456@(localhost:3306)/test?parseTime=True&loc=Local"

	db, err := gorm.Open(
		newDialector(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	replicaDSN := "root:123456@(localhost:3306)/test?parseTime=True&loc=Local"
	resolverConfig.Replicas = []gorm.Dialector{newDialector(replicaDSN)}

	err = db.Use(dbresolver.Register(resolverConfig).SetConnMaxLifetime(time.Minute * 10).SetMaxOpenConns(50).
		SetMaxIdleConns(15))
	if err != nil {
		panic(err)
	}

}

func newDialector(dsn string) gorm.Dialector {
	return mysql.Open(dsn)
}
