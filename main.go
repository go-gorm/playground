package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)
//
type User struct {
	gorm.Model
	Name          string
	Balance       int64 `json:"usdt_balance" gorm:"default:0; comment:金额"`
	UpdateVersion int64 `json:"update_version" gorm:"default:0;autoUpdateTime:nano;comment:纳秒时间戳,每次涉及到金钱的更新更新此字段,乐观锁冲突控制"`
}

type UpdateBalance struct {
	gorm.Model
	Name          string
	BeforeBalance int64 `json:"before_balance" gorm:"default:0; comment:金额变化前"`
	Balance       int64 `json:"balance" gorm:"default:0; comment:金额"`
	AfterBalance  int64 `json:"after_balance" gorm:"default:0; comment:金额变化后"`
}
func createMysql(address string) gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN:                       address, // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,     // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,    // disable datetime precision support, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,    // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,    // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,   // smart configure based on used version
	})
}
func main() {
	address := "root:88888888@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local"
	master := createMysql(address)
	db, err := gorm.Open(master)

	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	_ = db.AutoMigrate(&User{}, &UpdateBalance{})
	db.Exec("DELETE FROM update_balances")
	// 创建
	var user User
	if db.Where("id = ?", 1).First(&user).Error != nil {
		db.Create(&User{Name: "D42", Balance: 10000})
	}
	//
	//count_jian := 0
	//count_jia := 0
	// 进行并发操作
	ch := make(chan int, 100)
	var wg sync.WaitGroup
	// start 2 goroutines
	for i := 0; i < 99; i++ {
		wg.Add(1)
		go func(id int) {
			for {
				i, ok := <-ch
				if !ok {
					wg.Done()
					return
				}
				if i%2 == 0 {
					//count_jian++
					e := UpdateUser(db, -666)
					if e != nil {
						fmt.Println("出错：" + e.Error())
					}
					if err := db.First(&user, 1).Error; err != nil {
						fmt.Println("查询错误" + err.Error())
					}
					//  TODO 问题2. 后面这个对象的修改不能使用 db.Save
					//user.Name = fmt.Sprintf("大萨达%d", i)
					//db.Save(&user)
					// 也就是代理里面只能出现这种修改不能出现 db.Save
					if err := db.Model(&user).
						Update("name", fmt.Sprintf("大萨达%d", i)).Error; err != nil {
					}
					//db.Save(&user)
				} else {
					//count_jia++
					e := UpdateUser(db, +666)
					if e != nil {
						fmt.Println("出错：" + e.Error())
					}
				}
			}
		}(i)
	}
	for i := 0; i < 1000; i++ {
		ch <- i
	}
	//e := UpdateUser(db, -100)
	//if e != nil {
	//	fmt.Println("出错：" + e.Error())
	//}
	close(ch)
	wg.Wait()
	// 锁住指定 id 的 User 记录
	if err := db.First(&user, 1).Error; err != nil {
		fmt.Println("查询错误" + err.Error())
	}
	var jian, jia int64
	db.Model(&UpdateBalance{}).Where("balance > ?", 0).Pluck("SUM(balance)", &jia)
	db.Model(&UpdateBalance{}).Where("balance < ?", 0).Pluck("SUM(balance)", &jian)
	fmt.Printf("减了%d元，加了%d元", jia, jian)
	fmt.Printf("剩余：%d", user.Balance)
	var count int64
	db.Model(&UpdateBalance{}).Count(&count)
	fmt.Printf("总共记录了%d条", count)
}

func UpdateUser(db *gorm.DB, b int64) error {
	var user User
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 锁住指定 id 的 User 记录
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&user, 1).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 更新操作...
	before := user.Balance
	if err := tx.Model(&user).
		Update("name", "卫栖梧").
		Update("balance", gorm.Expr("balance + ?", b)).Error; err != nil {
		tx.Rollback()
		return err
	}

	after := user.Balance + b
	fmt.Println(before)
	fmt.Println(after)
	//  TODO  问题1.这里查出来的不对 新增记录就会出现问题
	up := UpdateBalance{
		Name:          user.Name,
		BeforeBalance: before,
		Balance:       b,
		AfterBalance:  after,
	}
	if err := tx.Create(&up).Error; err != nil {
		return err
	}

	//commit事务，释放锁
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
