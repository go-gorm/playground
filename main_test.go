package main

import (
	"net/http"
	_ "net/http/pprof"
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func leak() {
	user := User{Name: "jinzhu"}

	for i := 0; i < 10000; i++ {
		DB.FirstOrCreate(&user)
	}
}
func loop_forever() {
	for true {
		time.Sleep(time.Second)
	}
}
func TestGORM(t *testing.T) {
	go func() {
		t.Log("check the url: http://localhost:3344/debug/pprof")
		e := http.ListenAndServe("0.0.0.0:3344", nil)
		if e != nil {
			panic(e)
		}
	}()

	leak()

	loop_forever()
}
