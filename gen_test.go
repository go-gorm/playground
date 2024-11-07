/**
 * @Author: liuzhan (lz@ifreedom.top)
 * @Description: gen_test.go
 * @Version: 1.0.0
 * @Date: 2024/11/7 15:09
 */

package main

import (
	"context"
	"gorm.io/gen/examples/dal"
	"playground/dal/query"
	"testing"
)

func TestGEN(t *testing.T) {
	generate()
}

func TestPl(t *testing.T) {
	qUser := query.Use(dal.DB).User
	qUserExt := query.Use(dal.DB).UserExt
	qUser.WithContext(context.Background()).Preload(
		qUser.UserAccountRelationInfo,
		qUser.UserAccountRelationInfo.AccountInfo,
	).Find()

	qUserExt.WithContext(context.Background()).Preload(
		qUserExt.UserInfo,
		qUserExt.UserInfo.UserAccountRelationInfo,
		qUserExt.UserInfo.UserAccountRelationInfo.AccountInfo,
		qUserExt.UserInfo.UserAccountRelationInfo.AccountInfo.CompanyInfo,
	).Find()
}
