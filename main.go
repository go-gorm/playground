package main

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/playground/entity"
)

func main() {
	fmt.Println("vim-go")

	fmt.Printf("db:%v\n", GormDB)

	var errCreate error
	var companyProjectList []entity.CompanyProject
	if companyProjectList, errCreate = gorm.G[entity.CompanyProject](GormDB).Find(context.Background()); errCreate != nil {
		fmt.Println("errCreate:", errCreate)
		return
	}
	fmt.Println("find:", companyProjectList)

	name := "project111"
	department := "lab"
	company := "xxx"
	companyProject1 := entity.CompanyProject{
		Name:        &name,
		Description: &department,
		Company:     &company,
	}

	account := "admin111"
	ctx := context.WithValue(context.Background(), "account", account)
	if err := gorm.G[entity.CompanyProject](GormDB).Create(ctx, &companyProject1); err != nil {
		fmt.Println("create error:", err)
		return
	}

	nameUpdate := "project111 update"
	companyProject1.Name = &nameUpdate
	accountUpdate := "admin111 update"
	ctxUpdate := context.WithValue(context.Background(), "account", accountUpdate)
	rowsAffected, err := gorm.G[entity.CompanyProject](GormDB).Where("id = ?", companyProject1.ID).
		Omit("id", "created_at", "created_account", "deleted_at", "deleted_account", "bind_id").
		Updates(ctxUpdate, companyProject1)
	if err != nil {
		fmt.Println("updated failed err:", err)
		return
	}
	fmt.Println("rowsAffected:", rowsAffected)
	fmt.Printf("updated success")
}
