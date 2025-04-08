/**
 * @Author: liuzhan (lz@ifreedom.top)
 * @Description: gen_test.go
 * @Version: 1.0.0
 * @Date: 2024/11/7 15:09
 */

package main

import (
	"context"
	"playground/dal/query"
	"playground/model"
	"testing"
)

func TestGEN(t *testing.T) {
	generate()
}

func TestPl(t *testing.T) {
	ctx := context.Background()
	qUser := query.Use(db).User
	// 创建一个新的用户
	newUser := &model.User{
		Username: "test",
	}
	if err := qUser.WithContext(ctx).Create(newUser); err != nil {
		t.Error(err)
	}
	// 更新所有状态为0
	if _, err := qUser.WithContext(ctx).Where(qUser.Status.Gt(0)).Update(qUser.Status, 0); err != nil {
		t.Error(err)
	}

	// 查询当前用户状态
	user, err := qUser.WithContext(ctx).Debug().Where(qUser.ID.Eq(newUser.ID)).First()
	if err != nil {
		t.Error(err)
	}
	t.Logf("befer save status: %d", user.Status)

	// 修改username并使用 save 保存
	user.Username = "test2"
	if err = qUser.WithContext(ctx).Debug().Save(user); err != nil {
		t.Error(err)
	}

	// 查询最新的状态
	user, err = qUser.WithContext(ctx).Where(qUser.ID.Eq(newUser.ID)).First()
	if err != nil {
		t.Error(err)
	}
	t.Logf("after save status: %d", user.Status)
}
