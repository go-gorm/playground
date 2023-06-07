package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func errHandler1(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		m := User{Name: "thing"}
		errHandler2(db, m)(c)
	}
}

func errHandler2(db *gorm.DB, model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := db.Create(&model).Error; err != nil {
			log.Fatalln("db.Create error:", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

func noErrHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		model := User{Name: "thing"}
		if err := db.Create(&model).Error; err != nil {
			log.Fatalln("db.Create error:", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

func TestGORM(t *testing.T) {
	r := gin.Default()
	r.GET("/1", noErrHandler(DB))
	r.GET("/2", errHandler1(DB))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/1", nil)
	r.ServeHTTP(w, req)
	log.Println("Req 1 status:", w.Result().StatusCode)

	w2 := httptest.NewRecorder()
	req2 := httptest.NewRequest("GET", "/2", nil)
	r.ServeHTTP(w2, req2)
	log.Println("Req 2 status:", w.Result().StatusCode)
}
