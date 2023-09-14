package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: mysql
// some change
func TestGORM(t *testing.T) {
	Name := "rediis"
	Pwd := "12345"
	assert.NotNil(t, DB.Exec("CREATE USER ? IDENTIFIED BY ?", Name, Pwd).Error)

	cmd := fmt.Sprintf("CREATE USER '%s' IDENTIFIED BY '%s'", Name, Pwd)
	assert.Nil(t, DB.Exec(cmd).Error)

}
