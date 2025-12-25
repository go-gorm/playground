package main


import (
	"context"
	"fmt"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TblTaskLock struct {
	ID   string          `gorm:"not null;type:varchar(512);default:'';primaryKey;comment:主键"`
	Name string          `gorm:"not null;default:'';comment:名"`
	D    time.Time       `gorm:"not null;default:current_timestamp(3);comment:时间"`
}

var dbDefaultConn *gorm.DB
// 表名
func (t *TblTaskLock) TableName() string {
	return TblTaskLockName
}

func main(){

	var err error
	dsn:="xx:xx@tcp(xx:xxxx)/xxxx?autocommit=true&charset=utf8mb4,utf8&loc=Local&parseTime=true&writeTimeout=600s&readTimeout=3600s&timeout=10s"
	dbDefaultConn,err=gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, 
	}), &gorm.Config{})
	if err != nil{
		fmt.Println(err)
		return
	}
	var t = TblTaskLock{
		ID:   "1",
		Name: "1",
		D:    time.Date(2021,9,9,10,10,10,0,time.Now().Location()),
	}
	if err=dbDefaultConn.Clauses(clause.OnConflict{
		OnConstraint: "PRIMARY KEY",
		DoUpdates: clause.AssignmentColumns([]string{"name", "d"}),
	}).Create(&t).Error;err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("before:",t.D)
	t.Name="2"
	t.D=t.D.Add(100 * time.Second)
	if err=dbDefaultConn.Clauses(clause.OnConflict{
		OnConstraint: "PRIMARY KEY",
		DoUpdates: clause.AssignmentColumns([]string{"name", "d"}),

	}).Create(&t).Error;err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("expect:",t.D)

	if err= dbDefaultConn.Where("id = 1").First(&t).Error;err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("after:",t.D)
}
